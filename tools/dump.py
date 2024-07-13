
# -*- coding: utf-8 -*-

from google.protobuf.descriptor_pb2 import DescriptorProto, FieldDescriptorProto, FileDescriptorProto
from google.protobuf.internal.decoder import _DecodeVarint
from collections import OrderedDict, defaultdict
from os.path import dirname, realpath
from argparse import ArgumentParser
from itertools import groupby
from pathlib import Path
from shutil import which
from os import makedirs

__import__('sys').path.append(dirname(realpath(__file__)) + '/..')
#
#pip install protobuf google
#

# 定义一个有序字典用于存储提取器
extractors = OrderedDict()

# 定义一个装饰器用于注册提取器
def register_extractor(**kwargs):
    def register_extractor_decorate(func):
      extractors[kwargs['name']] = {'func': func, **kwargs}
      return func
    return register_extractor_decorate

def assert_installed(binaries=[]):
    missing = defaultdict(list)
    for item in binaries:
        if not which(item):
            missing["binaries"].append(item)

    if missing:
        msg = []
        for subject, names in missing.items():
            if len(names) == 1:
                subject = 'binary'
            msg.append('%s "%s"' % (subject, '", "'.join(names)))
        msg = 'You are missing the %s for this.' % ' and '.join(msg)
        print(msg) 

    return not missing

def extractor_save(base_path, folder, outputs):
    nb_written = 0
    name_to_path = {}
    wrote_endpoints = False
    
    for name, contents in outputs:
        if '.proto' in name:
            if folder:
                path = base_path / 'protos' / folder / name
            else:
                path = base_path / name
            
            makedirs(str(path.parent), exist_ok=True)
            with open(str(path), 'w') as fd:
                fd.write(contents)
            
            if name not in name_to_path:
                nb_written += 1
            name_to_path[name] = str(path)
        
        elif name.endswith('.sample'):
            endpoint = contents
            
            name = name.replace('.sample', '.proto')
            endpoint['proto_path'] = name_to_path[name]
            endpoint['proto_msg'] = name.replace('.proto', '')
    
    return nb_written, wrote_endpoints
#核心代码，其他的均为整理导出的proto  
def extractor_main(extractor):
    extractor = extractors[extractor]

    if assert_installed(**extractor.get('depends', {})):
        parser = ArgumentParser(description=extractor['desc'])
        parser.add_argument('input_', metavar='input_file')
        parser.add_argument('output_dir', type=Path, default='.', nargs='?')
        args = parser.parse_args()
        
        nb_written, wrote_endpoints = extractor_save(args.output_dir, '', extractor['func'](args.input_))
        if nb_written:
            print('\n[成功] 导出%s个proto文件到%s目录.\n' % (nb_written, args.output_dir))

INDENT = ' ' * 4

def descpb_to_proto(desc):
    out = 'syntax = "%s";\n\n' % (desc.syntax or 'proto2')

    scopes = ['']
    if desc.package:
        out += 'package %s;\n\n' % desc.package
        scopes[0] += '.' + desc.package
    
    for index, dep in enumerate(desc.dependency):
        prefix = ' public' * (index in desc.public_dependency)
        prefix += ' weak' * (index in desc.weak_dependency)
        out += 'import%s "%s";\n' % (prefix, dep)
        scopes.append('.' + ('/' + dep.rsplit('/', 1)[0])[1:].replace('/', '.'))
    
    out += '\n' * (out[-2] != '\n')
    
    out += parse_msg(desc, scopes, desc.syntax).strip('\n')
    name = desc.name.replace('..', '').strip('.\\/')
    
    return name, out + '\n'

def parse_msg(desc, scopes, syntax):
    out = ''
    is_msg = isinstance(desc, DescriptorProto)
    
    if is_msg:
        scopes = list(scopes)
        scopes[0] += '.' + desc.name
    
    blocks = OrderedDict()
    for nested_msg in (desc.nested_type if is_msg else desc.message_type):
        blocks[nested_msg.name] = parse_msg(nested_msg, scopes, syntax)
    
    for enum in desc.enum_type:
        out2 = ''
        for val in enum.value:
            out2 += '%s = %s;\n' % (val.name, fmt_value(val.number, val.options))
        
        if len(set(i.number for i in enum.value)) == len(enum.value):
            enum.options.ClearField('allow_alias')
        
        blocks[enum.name] = wrap_block('enum', out2, enum)
    
    if is_msg and desc.options.map_entry:
        return ' map<%s>' % ', '.join(min_name(i.type_name, scopes) \
            if i.type_name else types[i.type] \
                for i in desc.field)
    
    if is_msg:
        for field in desc.field:
            out += fmt_field(field, scopes, blocks, syntax)
        
        for index, oneof in enumerate(desc.oneof_decl):
            out += wrap_block('oneof', blocks.pop('_oneof_%d' % index), oneof)
        
        out += fmt_ranges('extensions', desc.extension_range)
        out += fmt_ranges('reserved', [*desc.reserved_range, *desc.reserved_name])
    
    else:
        for service in desc.service:
            out2 = ''
            for method in service.method:
                out2 += 'rpc %s(%s%s) returns (%s%s);\n' % (method.name,
                    'stream ' * method.client_streaming,
                    min_name(method.input_type, scopes),
                    'stream ' * method.server_streaming,
                    min_name(method.output_type, scopes))
            
            out += wrap_block('service', out2, service)
    
    extendees = OrderedDict()
    for ext in desc.extension:
        extendees.setdefault(ext.extendee, '')
        extendees[ext.extendee] += fmt_field(ext, scopes, blocks, syntax, True)
    
    for name, value in blocks.items():
        out += value[:-1]
    
    for name, fields in extendees.items():
        out += wrap_block('extend', fields, name=min_name(name, scopes))
    
    out = wrap_block('message' * is_msg, out, desc)
    return out

def fmt_value(val, options=None, desc=None, optarr=[]):
    if type(val) != str:
        if type(val) == bool:
            val = str(val).lower()
        elif desc and desc.enum_type:
            val = desc.enum_type.values_by_number[val].name
        val = str(val)
    else:
        val = '"%s"' % val.encode('unicode_escape').decode('utf8')
    
    if options:
        opts = [*optarr]
        for (option, value) in options.ListFields():
            opts.append('%s = %s' % (option.name, fmt_value(value, desc=option)))
        if opts:
            val += ' [%s]' % ', '.join(opts)
    return val

types = {v: k.split('_')[1].lower() for k, v in FieldDescriptorProto.Type.items()}
labels = {v: k.split('_')[1].lower() for k, v in FieldDescriptorProto.Label.items()}

def fmt_field(field, scopes, blocks, syntax, extend=False):
    type_ = types[field.type]
    
    default = ''
    if field.default_value:
        if field.type == field.TYPE_STRING:
            default = ['default = %s' % fmt_value(field.default_value)]
        elif field.type == field.TYPE_BYTES:
            default = ['default = "%s"' % field.default_value]
        else:
            
            if ('int' in type_ or 'fixed' in type_) and \
               int(field.default_value) >= 0x10000 and \
               not any(len(list(i)) > 3 for _, i in groupby(str(field.default_value))):
                
                field.default_value = hex(int(field.default_value))
            
            default = ['default = %s' % field.default_value]
    
    out = ''
    if field.type_name:
        type_ = min_name(field.type_name, scopes)
        short_type = type_.split('.')[-1]
        
        if short_type in blocks and ((not extend and not field.HasField('oneof_index')) or \
                                      blocks[short_type].startswith(' map<')):
            out += blocks.pop(short_type)[1:]
    
    if out.startswith('map<'):
        line = out + ' %s = %s;\n' % (field.name, fmt_value(field.number, field.options, optarr=default))
        out = ''
    elif field.type != field.TYPE_GROUP:
        line = '%s %s %s = %s;\n' % (labels[field.label], type_, field.name, fmt_value(field.number, field.options, optarr=default))
    else:
        line = '%s group %s = %d ' % (labels[field.label], type_, field.number)
        out = out.split(' ', 2)[-1]
    
    if field.HasField('oneof_index') or (syntax == 'proto3' and line.startswith('optional')):
        line = line.split(' ', 1)[-1]
    if out:
        line = '\n' + line
    
    if field.HasField('oneof_index'):
        blocks.setdefault('_oneof_%d' % field.oneof_index, '')
        blocks['_oneof_%d' % field.oneof_index] += line + out
        return ''
    else:
        return line + out

def min_name(name, scopes):
    name, cur_scope = name.split('.'), scopes[0].split('.')
    short_name = [name.pop()]
    
    while name and (cur_scope[:len(name)] != name or \
                    any(list_rfind(scope.split('.'), short_name[0]) > len(name) \
                        for scope in scopes)):
        short_name.insert(0, name.pop())
    
    return '.'.join(short_name)

def wrap_block(type_, value, desc=None, name=None):
    out = ''
    if type_:
        out = '\n%s %s {\n' % (type_, name or desc.name)
    
    if desc:
        for (option, optval) in desc.options.ListFields():
            value = 'option %s = %s;\n' % (option.name, fmt_value(optval, desc=option)) + value
    
    value = value.replace('\n\n\n', '\n\n')
    if type_:
        out += '\n'.join(INDENT + line for line in value.strip('\n').split('\n'))
        out += '\n}\n\n'
    else:
        out += value
    return out

def fmt_ranges(name, ranges):
    text = []
    for range_ in ranges:
        if type(range_) != str and range_.end - 1 > range_.start:
            if range_.end < 0x20000000:
                text.append('%d to %d' % (range_.start, range_.end - 1))
            else:
                text.append('%d to max' % range_.start)
        elif type(range_) != str:
            text.append(fmt_value(range_.start))
        else:
            text.append(fmt_value(range_))
    if text:
        return '\n%s %s;\n' % (name, ', '.join(text))
    return ''

list_rfind = lambda x, i: len(x) - 1 - x[::-1].index(i) if i in x else -1

@register_extractor(name = 'from_binary',
                    desc = 'Extract Protobuf metadata from binary file (*.dll, *.so...)')
#核心代码，其他的均为整理导出的proto  
def walk_binary(binr):
    if type(binr) == str:
        try:
            with open(binr, 'rb') as fd:
                binr = fd.read()
        except Exception:
            return
    
    cursor = 0
    while cursor < len(binr):
        cursor = binr.find(b'.proto', cursor)
        
        if cursor == -1:
            break
        cursor += len('.proto')
        cursor += (binr[cursor:cursor + 5] == b'devel') * 5
        
        start = binr.rfind(b'\x0a', max(cursor - 1024, 0), cursor)
        
        if start > 0 and binr[start - 1] == 0x0a == (cursor - start - 1):
            start -= 1
        
        if start == -1:
            continue
        varint, end = _DecodeVarint(binr, start + 1)
        if cursor - end != varint:
            continue
        
        tags = b'\x12\x1a\x22\x2a\x32\x3a\x42\x4a\x50\x58\x62'
        if binr[cursor] not in tags:
            continue
        
        while cursor < len(binr) and binr[cursor] in tags:
            tags = tags[tags.index(binr[cursor]):]
            
            varint, end = _DecodeVarint(binr, cursor + 1)
            cursor = end + varint * (binr[cursor] & 0b111 == 2)
        
        proto = FileDescriptorProto()
        proto.ParseFromString(binr[start:cursor])
        
        yield descpb_to_proto(proto)


if __name__ == '__main__':
    extractor_main('from_binary')