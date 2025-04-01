#
# Autogenerated by Thrift
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#  @generated
#

from __future__ import absolute_import
import sys
from thrift.util.Recursive import fix_spec
from thrift.Thrift import TType, TMessageType, TPriority, TRequestContext, TProcessorEventHandler, TServerInterface, TProcessor, TException, TApplicationException, UnimplementedTypedef
from thrift.protocol.TProtocol import TProtocolException

from json import loads
import sys
if sys.version_info[0] >= 3:
  long = int


import pprint
import warnings
from thrift import Thrift
from thrift.transport import TTransport
from thrift.protocol import TBinaryProtocol
from thrift.protocol import TCompactProtocol
from thrift.protocol import THeaderProtocol
fastproto = None
try:
  from thrift.protocol import fastproto
except ImportError:
  pass

def __EXPAND_THRIFT_SPEC(spec):
    next_id = 0
    for item in spec:
        item_id = item[0]
        if next_id >= 0 and item_id < 0:
            next_id = item_id
        if item_id != next_id:
            for _ in range(next_id, item_id):
                yield None
        yield item
        next_id = item_id + 1

class ThriftEnumWrapper(int):
  def __new__(cls, enum_class, value):
    return super().__new__(cls, value)
  def __init__(self, enum_class, value):    self.enum_class = enum_class
  def __repr__(self):
    return self.enum_class.__name__ + '.' + self.enum_class._VALUES_TO_NAMES[self]

all_structs = []
UTF8STRINGS = bool(0) or sys.version_info.major >= 3

__all__ = ['UTF8STRINGS', 'MyEnum', 'MyStructNestedAnnotation', 'MyUnion', 'MyException', 'MyStruct', 'SecretStruct', 'AwesomeStruct', 'FantasticStruct', 'list_string_6884']

class MyEnum:

  _NAMES_TO_VALUES = dict(zip((
    "MyValue1",
    "MyValue2",
    "DOMAIN",
),
(
    0,
    1,
    2,
  )))
  _VALUES_TO_NAMES = {}

for k, v in MyEnum._NAMES_TO_VALUES.items():
    setattr(MyEnum, k, v)
    MyEnum._VALUES_TO_NAMES[v] = k

class MyStructNestedAnnotation:
  r"""
  Attributes:
   - name
  """

  thrift_spec = None
  thrift_field_annotations = None
  thrift_struct_annotations = None
  __init__ = None
  @staticmethod
  def isUnion():
    return False

  def read(self, iprot):
    if (isinstance(iprot, TBinaryProtocol.TBinaryProtocolAccelerated) or (isinstance(iprot, THeaderProtocol.THeaderProtocolAccelerate) and iprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_BINARY_PROTOCOL)) and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None and fastproto is not None:
      fastproto.decode(self, iprot.trans, [self.__class__, self.thrift_spec, False], utf8strings=UTF8STRINGS, protoid=0)
      return
    if (isinstance(iprot, TCompactProtocol.TCompactProtocolAccelerated) or (isinstance(iprot, THeaderProtocol.THeaderProtocolAccelerate) and iprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_COMPACT_PROTOCOL)) and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None and fastproto is not None:
      fastproto.decode(self, iprot.trans, [self.__class__, self.thrift_spec, False], utf8strings=UTF8STRINGS, protoid=2)
      return
    iprot.readStructBegin()
    while True:
      (fname, ftype, fid) = iprot.readFieldBegin()
      if ftype == TType.STOP:
        break
      if fid == 1:
        if ftype == TType.STRING:
          self.name = iprot.readString().decode('utf-8') if UTF8STRINGS else iprot.readString()
        else:
          iprot.skip(ftype)
      else:
        iprot.skip(ftype)
      iprot.readFieldEnd()
    iprot.readStructEnd()

  def write(self, oprot):
    if (isinstance(oprot, TBinaryProtocol.TBinaryProtocolAccelerated) or (isinstance(oprot, THeaderProtocol.THeaderProtocolAccelerate) and oprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_BINARY_PROTOCOL)) and self.thrift_spec is not None and fastproto is not None:
      oprot.trans.write(fastproto.encode(self, [self.__class__, self.thrift_spec, False], utf8strings=UTF8STRINGS, protoid=0))
      return
    if (isinstance(oprot, TCompactProtocol.TCompactProtocolAccelerated) or (isinstance(oprot, THeaderProtocol.THeaderProtocolAccelerate) and oprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_COMPACT_PROTOCOL)) and self.thrift_spec is not None and fastproto is not None:
      oprot.trans.write(fastproto.encode(self, [self.__class__, self.thrift_spec, False], utf8strings=UTF8STRINGS, protoid=2))
      return
    oprot.writeStructBegin('MyStructNestedAnnotation')
    if self.name != None:
      oprot.writeFieldBegin('name', TType.STRING, 1)
      oprot.writeString(self.name.encode('utf-8')) if UTF8STRINGS and not isinstance(self.name, bytes) else oprot.writeString(self.name)
      oprot.writeFieldEnd()
    oprot.writeFieldStop()
    oprot.writeStructEnd()

  def readFromJson(self, json, is_text=True, **kwargs):
    kwargs_copy = dict(kwargs)
    relax_enum_validation = bool(kwargs_copy.pop('relax_enum_validation', False))
    set_cls = kwargs_copy.pop('custom_set_cls', set)
    dict_cls = kwargs_copy.pop('custom_dict_cls', dict)
    wrap_enum_constants = kwargs_copy.pop('wrap_enum_constants', False)
    if wrap_enum_constants and relax_enum_validation:
        raise ValueError(
            'wrap_enum_constants cannot be used together with relax_enum_validation'
        )
    if kwargs_copy:
        extra_kwargs = ', '.join(kwargs_copy.keys())
        raise ValueError(
            'Unexpected keyword arguments: ' + extra_kwargs
        )
    json_obj = json
    if is_text:
      json_obj = loads(json)
    if 'name' in json_obj and json_obj['name'] is not None:
      self.name = json_obj['name']

  def __repr__(self):
    L = []
    padding = ' ' * 4
    if self.name is not None:
      value = pprint.pformat(self.name, indent=0)
      value = padding.join(value.splitlines(True))
      L.append('    name=%s' % (value))
    return "%s(%s)" % (self.__class__.__name__, "\n" + ",\n".join(L) if L else '')

  def __eq__(self, other):
    if not isinstance(other, self.__class__):
      return False

    return self.__dict__ == other.__dict__ 

  def __ne__(self, other):
    return not (self == other)

  def __dir__(self):
    return (
      'name',
    )

  __hash__ = object.__hash__

  def _to_python(self):
    import importlib
    import thrift.python.converter
    python_types = importlib.import_module("module.thrift_types")
    return thrift.python.converter.to_python_struct(python_types.MyStructNestedAnnotation, self)

  def _to_mutable_python(self):
    import importlib
    import thrift.python.mutable_converter
    python_mutable_types = importlib.import_module("module.thrift_mutable_types")
    return thrift.python.mutable_converter.to_mutable_python_struct_or_union(python_mutable_types.MyStructNestedAnnotation, self)

  def _to_py3(self):
    import importlib
    import thrift.py3.converter
    py3_types = importlib.import_module("module.types")
    return thrift.py3.converter.to_py3_struct(py3_types.MyStructNestedAnnotation, self)

  def _to_py_deprecated(self):
    return self

class MyUnion(object):

  thrift_spec = None
  __EMPTY__ = 0
  
  @staticmethod
  def isUnion():
    return True

  def getType(self):
    return self.field

  def __repr__(self):
    value = pprint.pformat(self.value)
    member = ''
    return "%s(%s)" % (self.__class__.__name__, member)

  def read(self, iprot):
    self.field = 0
    self.value = None
    if (isinstance(iprot, TBinaryProtocol.TBinaryProtocolAccelerated) or (isinstance(iprot, THeaderProtocol.THeaderProtocolAccelerate) and iprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_BINARY_PROTOCOL)) and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None and fastproto is not None:
      fastproto.decode(self, iprot.trans, [self.__class__, self.thrift_spec, True], utf8strings=UTF8STRINGS, protoid=0)
      return
    if (isinstance(iprot, TCompactProtocol.TCompactProtocolAccelerated) or (isinstance(iprot, THeaderProtocol.THeaderProtocolAccelerate) and iprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_COMPACT_PROTOCOL)) and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None and fastproto is not None:
      fastproto.decode(self, iprot.trans, [self.__class__, self.thrift_spec, True], utf8strings=UTF8STRINGS, protoid=2)
      return
    iprot.readStructBegin()
    while True:
      (fname, ftype, fid) = iprot.readFieldBegin()
      if ftype == TType.STOP:
        break

      else:
        iprot.skip(ftype)
      iprot.readFieldEnd()
    iprot.readStructEnd()

  def write(self, oprot):
    if (isinstance(oprot, TBinaryProtocol.TBinaryProtocolAccelerated) or (isinstance(oprot, THeaderProtocol.THeaderProtocolAccelerate) and oprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_BINARY_PROTOCOL)) and self.thrift_spec is not None and fastproto is not None:
      oprot.trans.write(fastproto.encode(self, [self.__class__, self.thrift_spec, True], utf8strings=UTF8STRINGS, protoid=0))
      return
    if (isinstance(oprot, TCompactProtocol.TCompactProtocolAccelerated) or (isinstance(oprot, THeaderProtocol.THeaderProtocolAccelerate) and oprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_COMPACT_PROTOCOL)) and self.thrift_spec is not None and fastproto is not None:
      oprot.trans.write(fastproto.encode(self, [self.__class__, self.thrift_spec, True], utf8strings=UTF8STRINGS, protoid=2))
      return
    oprot.writeUnionBegin('MyUnion')
    oprot.writeFieldStop()
    oprot.writeUnionEnd()
  
  def readFromJson(self, json, is_text=True, **kwargs):
    kwargs_copy = dict(kwargs)
    relax_enum_validation = bool(kwargs_copy.pop('relax_enum_validation', False))
    set_cls = kwargs_copy.pop('custom_set_cls', set)
    dict_cls = kwargs_copy.pop('custom_dict_cls', dict)
    wrap_enum_constants = kwargs_copy.pop('wrap_enum_constants', False)
    if wrap_enum_constants and relax_enum_validation:
        raise ValueError(
            'wrap_enum_constants cannot be used together with relax_enum_validation'
        )
    if kwargs_copy:
        extra_kwargs = ', '.join(kwargs_copy.keys())
        raise ValueError(
            'Unexpected keyword arguments: ' + extra_kwargs
        )
    self.field = 0
    self.value = None
    obj = json
    if is_text:
      obj = loads(json)
    if not isinstance(obj, dict) or len(obj) > 1:
      raise TProtocolException(TProtocolException.INVALID_DATA, 'Can not parse')
    

  def __eq__(self, other):
    if not isinstance(other, self.__class__):
      return False

    return self.__dict__ == other.__dict__

  def __ne__(self, other):
    return not (self == other)

  def _to_python(self):
    import importlib
    import thrift.python.converter
    python_types = importlib.import_module("module.thrift_types")
    return thrift.python.converter.to_python_struct(python_types.MyUnion, self)

  def _to_mutable_python(self):
    import importlib
    import thrift.python.mutable_converter
    python_mutable_types = importlib.import_module("module.thrift_mutable_types")
    return thrift.python.mutable_converter.to_mutable_python_struct_or_union(python_mutable_types.MyUnion, self)

  def _to_py3(self):
    import importlib
    import thrift.py3.converter
    py3_types = importlib.import_module("module.types")
    return thrift.py3.converter.to_py3_struct(py3_types.MyUnion, self)

  def _to_py_deprecated(self):
    return self

class MyException(TException):

  thrift_spec = None
  thrift_field_annotations = None
  thrift_struct_annotations = None
  @staticmethod
  def isUnion():
    return False

  def read(self, iprot):
    if (isinstance(iprot, TBinaryProtocol.TBinaryProtocolAccelerated) or (isinstance(iprot, THeaderProtocol.THeaderProtocolAccelerate) and iprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_BINARY_PROTOCOL)) and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None and fastproto is not None:
      fastproto.decode(self, iprot.trans, [self.__class__, self.thrift_spec, False], utf8strings=UTF8STRINGS, protoid=0)
      return
    if (isinstance(iprot, TCompactProtocol.TCompactProtocolAccelerated) or (isinstance(iprot, THeaderProtocol.THeaderProtocolAccelerate) and iprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_COMPACT_PROTOCOL)) and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None and fastproto is not None:
      fastproto.decode(self, iprot.trans, [self.__class__, self.thrift_spec, False], utf8strings=UTF8STRINGS, protoid=2)
      return
    iprot.readStructBegin()
    while True:
      (fname, ftype, fid) = iprot.readFieldBegin()
      if ftype == TType.STOP:
        break
      else:
        iprot.skip(ftype)
      iprot.readFieldEnd()
    iprot.readStructEnd()

  def write(self, oprot):
    if (isinstance(oprot, TBinaryProtocol.TBinaryProtocolAccelerated) or (isinstance(oprot, THeaderProtocol.THeaderProtocolAccelerate) and oprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_BINARY_PROTOCOL)) and self.thrift_spec is not None and fastproto is not None:
      oprot.trans.write(fastproto.encode(self, [self.__class__, self.thrift_spec, False], utf8strings=UTF8STRINGS, protoid=0))
      return
    if (isinstance(oprot, TCompactProtocol.TCompactProtocolAccelerated) or (isinstance(oprot, THeaderProtocol.THeaderProtocolAccelerate) and oprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_COMPACT_PROTOCOL)) and self.thrift_spec is not None and fastproto is not None:
      oprot.trans.write(fastproto.encode(self, [self.__class__, self.thrift_spec, False], utf8strings=UTF8STRINGS, protoid=2))
      return
    oprot.writeStructBegin('MyException')
    oprot.writeFieldStop()
    oprot.writeStructEnd()

  def readFromJson(self, json, is_text=True, **kwargs):
    kwargs_copy = dict(kwargs)
    relax_enum_validation = bool(kwargs_copy.pop('relax_enum_validation', False))
    set_cls = kwargs_copy.pop('custom_set_cls', set)
    dict_cls = kwargs_copy.pop('custom_dict_cls', dict)
    wrap_enum_constants = kwargs_copy.pop('wrap_enum_constants', False)
    if wrap_enum_constants and relax_enum_validation:
        raise ValueError(
            'wrap_enum_constants cannot be used together with relax_enum_validation'
        )
    if kwargs_copy:
        extra_kwargs = ', '.join(kwargs_copy.keys())
        raise ValueError(
            'Unexpected keyword arguments: ' + extra_kwargs
        )
    json_obj = json
    if is_text:
      json_obj = loads(json)

  def __str__(self):
    return repr(self)

  def __repr__(self):
    L = []
    padding = ' ' * 4
    if 'message' not in self.__dict__:
      message = getattr(self, 'message', None)
      if message:
        L.append('message=%r' % message)
    return "%s(%s)" % (self.__class__.__name__, "\n" + ",\n".join(L) if L else '')

  def __eq__(self, other):
    if not isinstance(other, self.__class__):
      return False

    return self.__dict__ == other.__dict__ 

  def __ne__(self, other):
    return not (self == other)

  def __dir__(self):
    return (
    )

  __hash__ = object.__hash__

  def _to_python(self):
    import importlib
    import thrift.python.converter
    python_types = importlib.import_module("module.thrift_types")
    return thrift.python.converter.to_python_struct(python_types.MyException, self)

  def _to_mutable_python(self):
    import importlib
    import thrift.python.mutable_converter
    python_mutable_types = importlib.import_module("module.thrift_mutable_types")
    return thrift.python.mutable_converter.to_mutable_python_struct_or_union(python_mutable_types.MyException, self)

  def _to_py3(self):
    import importlib
    import thrift.py3.converter
    py3_types = importlib.import_module("module.types")
    return thrift.py3.converter.to_py3_struct(py3_types.MyException, self)

  def _to_py_deprecated(self):
    return self

class MyStruct:
  r"""
  Attributes:
   - major
   - abstract
   - annotation_with_quote
   - class_
   - annotation_with_trailing_comma
   - empty_annotations
   - my_enum
   - cpp_type_annotation
   - my_union
  """

  thrift_spec = None
  thrift_field_annotations = None
  thrift_struct_annotations = None
  __init__ = None
  @staticmethod
  def isUnion():
    return False

  def read(self, iprot):
    if (isinstance(iprot, TBinaryProtocol.TBinaryProtocolAccelerated) or (isinstance(iprot, THeaderProtocol.THeaderProtocolAccelerate) and iprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_BINARY_PROTOCOL)) and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None and fastproto is not None:
      fastproto.decode(self, iprot.trans, [self.__class__, self.thrift_spec, False], utf8strings=UTF8STRINGS, protoid=0)
      return
    if (isinstance(iprot, TCompactProtocol.TCompactProtocolAccelerated) or (isinstance(iprot, THeaderProtocol.THeaderProtocolAccelerate) and iprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_COMPACT_PROTOCOL)) and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None and fastproto is not None:
      fastproto.decode(self, iprot.trans, [self.__class__, self.thrift_spec, False], utf8strings=UTF8STRINGS, protoid=2)
      return
    iprot.readStructBegin()
    while True:
      (fname, ftype, fid) = iprot.readFieldBegin()
      if ftype == TType.STOP:
        break
      if fid == 2:
        if ftype == TType.I64:
          self.major = iprot.readI64()
        else:
          iprot.skip(ftype)
      elif fid == 1:
        if ftype == TType.STRING:
          self.abstract = iprot.readString().decode('utf-8') if UTF8STRINGS else iprot.readString()
        else:
          iprot.skip(ftype)
      elif fid == 3:
        if ftype == TType.STRING:
          self.annotation_with_quote = iprot.readString().decode('utf-8') if UTF8STRINGS else iprot.readString()
        else:
          iprot.skip(ftype)
      elif fid == 4:
        if ftype == TType.STRING:
          self.class_ = iprot.readString().decode('utf-8') if UTF8STRINGS else iprot.readString()
        else:
          iprot.skip(ftype)
      elif fid == 5:
        if ftype == TType.STRING:
          self.annotation_with_trailing_comma = iprot.readString().decode('utf-8') if UTF8STRINGS else iprot.readString()
        else:
          iprot.skip(ftype)
      elif fid == 6:
        if ftype == TType.STRING:
          self.empty_annotations = iprot.readString().decode('utf-8') if UTF8STRINGS else iprot.readString()
        else:
          iprot.skip(ftype)
      elif fid == 7:
        if ftype == TType.I32:
          self.my_enum = iprot.readI32()
        else:
          iprot.skip(ftype)
      elif fid == 8:
        if ftype == TType.LIST:
          self.cpp_type_annotation = []
          (_etype3, _size0) = iprot.readListBegin()
          if _size0 >= 0:
            for _i4 in range(_size0):
              _elem5 = iprot.readString().decode('utf-8') if UTF8STRINGS else iprot.readString()
              self.cpp_type_annotation.append(_elem5)
          else: 
            while iprot.peekList():
              _elem6 = iprot.readString().decode('utf-8') if UTF8STRINGS else iprot.readString()
              self.cpp_type_annotation.append(_elem6)
          iprot.readListEnd()
        else:
          iprot.skip(ftype)
      elif fid == 9:
        if ftype == TType.STRUCT:
          self.my_union = MyUnion()
          self.my_union.read(iprot)
        else:
          iprot.skip(ftype)
      else:
        iprot.skip(ftype)
      iprot.readFieldEnd()
    iprot.readStructEnd()

  def write(self, oprot):
    if (isinstance(oprot, TBinaryProtocol.TBinaryProtocolAccelerated) or (isinstance(oprot, THeaderProtocol.THeaderProtocolAccelerate) and oprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_BINARY_PROTOCOL)) and self.thrift_spec is not None and fastproto is not None:
      oprot.trans.write(fastproto.encode(self, [self.__class__, self.thrift_spec, False], utf8strings=UTF8STRINGS, protoid=0))
      return
    if (isinstance(oprot, TCompactProtocol.TCompactProtocolAccelerated) or (isinstance(oprot, THeaderProtocol.THeaderProtocolAccelerate) and oprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_COMPACT_PROTOCOL)) and self.thrift_spec is not None and fastproto is not None:
      oprot.trans.write(fastproto.encode(self, [self.__class__, self.thrift_spec, False], utf8strings=UTF8STRINGS, protoid=2))
      return
    oprot.writeStructBegin('MyStruct')
    if self.abstract != None:
      oprot.writeFieldBegin('abstract', TType.STRING, 1)
      oprot.writeString(self.abstract.encode('utf-8')) if UTF8STRINGS and not isinstance(self.abstract, bytes) else oprot.writeString(self.abstract)
      oprot.writeFieldEnd()
    if self.major != None:
      oprot.writeFieldBegin('major', TType.I64, 2)
      oprot.writeI64(self.major)
      oprot.writeFieldEnd()
    if self.annotation_with_quote != None:
      oprot.writeFieldBegin('annotation_with_quote', TType.STRING, 3)
      oprot.writeString(self.annotation_with_quote.encode('utf-8')) if UTF8STRINGS and not isinstance(self.annotation_with_quote, bytes) else oprot.writeString(self.annotation_with_quote)
      oprot.writeFieldEnd()
    if self.class_ != None:
      oprot.writeFieldBegin('class_', TType.STRING, 4)
      oprot.writeString(self.class_.encode('utf-8')) if UTF8STRINGS and not isinstance(self.class_, bytes) else oprot.writeString(self.class_)
      oprot.writeFieldEnd()
    if self.annotation_with_trailing_comma != None:
      oprot.writeFieldBegin('annotation_with_trailing_comma', TType.STRING, 5)
      oprot.writeString(self.annotation_with_trailing_comma.encode('utf-8')) if UTF8STRINGS and not isinstance(self.annotation_with_trailing_comma, bytes) else oprot.writeString(self.annotation_with_trailing_comma)
      oprot.writeFieldEnd()
    if self.empty_annotations != None:
      oprot.writeFieldBegin('empty_annotations', TType.STRING, 6)
      oprot.writeString(self.empty_annotations.encode('utf-8')) if UTF8STRINGS and not isinstance(self.empty_annotations, bytes) else oprot.writeString(self.empty_annotations)
      oprot.writeFieldEnd()
    if self.my_enum != None:
      oprot.writeFieldBegin('my_enum', TType.I32, 7)
      oprot.writeI32(self.my_enum)
      oprot.writeFieldEnd()
    if self.cpp_type_annotation != None:
      oprot.writeFieldBegin('cpp_type_annotation', TType.LIST, 8)
      oprot.writeListBegin(TType.STRING, len(self.cpp_type_annotation))
      for iter7 in self.cpp_type_annotation:
        oprot.writeString(iter7.encode('utf-8')) if UTF8STRINGS and not isinstance(iter7, bytes) else oprot.writeString(iter7)
      oprot.writeListEnd()
      oprot.writeFieldEnd()
    if self.my_union != None:
      oprot.writeFieldBegin('my_union', TType.STRUCT, 9)
      self.my_union.write(oprot)
      oprot.writeFieldEnd()
    oprot.writeFieldStop()
    oprot.writeStructEnd()

  def readFromJson(self, json, is_text=True, **kwargs):
    kwargs_copy = dict(kwargs)
    relax_enum_validation = bool(kwargs_copy.pop('relax_enum_validation', False))
    set_cls = kwargs_copy.pop('custom_set_cls', set)
    dict_cls = kwargs_copy.pop('custom_dict_cls', dict)
    wrap_enum_constants = kwargs_copy.pop('wrap_enum_constants', False)
    if wrap_enum_constants and relax_enum_validation:
        raise ValueError(
            'wrap_enum_constants cannot be used together with relax_enum_validation'
        )
    if kwargs_copy:
        extra_kwargs = ', '.join(kwargs_copy.keys())
        raise ValueError(
            'Unexpected keyword arguments: ' + extra_kwargs
        )
    json_obj = json
    if is_text:
      json_obj = loads(json)
    if 'major' in json_obj and json_obj['major'] is not None:
      self.major = long(json_obj['major'])
    if 'abstract' in json_obj and json_obj['abstract'] is not None:
      self.abstract = json_obj['abstract']
    if 'annotation_with_quote' in json_obj and json_obj['annotation_with_quote'] is not None:
      self.annotation_with_quote = json_obj['annotation_with_quote']
    if 'class_' in json_obj and json_obj['class_'] is not None:
      self.class_ = json_obj['class_']
    if 'annotation_with_trailing_comma' in json_obj and json_obj['annotation_with_trailing_comma'] is not None:
      self.annotation_with_trailing_comma = json_obj['annotation_with_trailing_comma']
    if 'empty_annotations' in json_obj and json_obj['empty_annotations'] is not None:
      self.empty_annotations = json_obj['empty_annotations']
    if 'my_enum' in json_obj and json_obj['my_enum'] is not None:
      self.my_enum = json_obj['my_enum']
      if not self.my_enum in MyEnum._VALUES_TO_NAMES:
        msg = 'Integer value ''%s'' is not a recognized value of enum type MyEnum' % self.my_enum
        if relax_enum_validation:
            warnings.warn(msg)
        else:
            raise TProtocolException(TProtocolException.INVALID_DATA, msg)
      if wrap_enum_constants:
        self.my_enum = ThriftEnumWrapper(MyEnum, self.my_enum)
    if 'cpp_type_annotation' in json_obj and json_obj['cpp_type_annotation'] is not None:
      self.cpp_type_annotation = []
      for _tmp_e8 in json_obj['cpp_type_annotation']:
        self.cpp_type_annotation.append(_tmp_e8)
    if 'my_union' in json_obj and json_obj['my_union'] is not None:
      self.my_union = MyUnion()
      self.my_union.readFromJson(json_obj['my_union'], is_text=False, **kwargs)

  def __repr__(self):
    L = []
    padding = ' ' * 4
    if self.major is not None:
      value = pprint.pformat(self.major, indent=0)
      value = padding.join(value.splitlines(True))
      L.append('    major=%s' % (value))
    if self.abstract is not None:
      value = pprint.pformat(self.abstract, indent=0)
      value = padding.join(value.splitlines(True))
      L.append('    abstract=%s' % (value))
    if self.annotation_with_quote is not None:
      value = pprint.pformat(self.annotation_with_quote, indent=0)
      value = padding.join(value.splitlines(True))
      L.append('    annotation_with_quote=%s' % (value))
    if self.class_ is not None:
      value = pprint.pformat(self.class_, indent=0)
      value = padding.join(value.splitlines(True))
      L.append('    class_=%s' % (value))
    if self.annotation_with_trailing_comma is not None:
      value = pprint.pformat(self.annotation_with_trailing_comma, indent=0)
      value = padding.join(value.splitlines(True))
      L.append('    annotation_with_trailing_comma=%s' % (value))
    if self.empty_annotations is not None:
      value = pprint.pformat(self.empty_annotations, indent=0)
      value = padding.join(value.splitlines(True))
      L.append('    empty_annotations=%s' % (value))
    if self.my_enum is not None:
      value = pprint.pformat(self.my_enum, indent=0)
      value = padding.join(value.splitlines(True))
      L.append('    my_enum=%s' % (value))
    if self.cpp_type_annotation is not None:
      value = pprint.pformat(self.cpp_type_annotation, indent=0)
      value = padding.join(value.splitlines(True))
      L.append('    cpp_type_annotation=%s' % (value))
    if self.my_union is not None:
      value = pprint.pformat(self.my_union, indent=0)
      value = padding.join(value.splitlines(True))
      L.append('    my_union=%s' % (value))
    return "%s(%s)" % (self.__class__.__name__, "\n" + ",\n".join(L) if L else '')

  def __eq__(self, other):
    if not isinstance(other, self.__class__):
      return False

    return self.__dict__ == other.__dict__ 

  def __ne__(self, other):
    return not (self == other)

  def __dir__(self):
    return (
      'abstract',
      'major',
      'annotation_with_quote',
      'class_',
      'annotation_with_trailing_comma',
      'empty_annotations',
      'my_enum',
      'cpp_type_annotation',
      'my_union',
    )

  __hash__ = object.__hash__

  def _to_python(self):
    import importlib
    import thrift.python.converter
    python_types = importlib.import_module("module.thrift_types")
    return thrift.python.converter.to_python_struct(python_types.MyStruct, self)

  def _to_mutable_python(self):
    import importlib
    import thrift.python.mutable_converter
    python_mutable_types = importlib.import_module("module.thrift_mutable_types")
    return thrift.python.mutable_converter.to_mutable_python_struct_or_union(python_mutable_types.MyStruct, self)

  def _to_py3(self):
    import importlib
    import thrift.py3.converter
    py3_types = importlib.import_module("module.types")
    return thrift.py3.converter.to_py3_struct(py3_types.MyStruct, self)

  def _to_py_deprecated(self):
    return self

class SecretStruct:
  r"""
  Attributes:
   - id
   - password
  """

  thrift_spec = None
  thrift_field_annotations = None
  thrift_struct_annotations = None
  __init__ = None
  @staticmethod
  def isUnion():
    return False

  def read(self, iprot):
    if (isinstance(iprot, TBinaryProtocol.TBinaryProtocolAccelerated) or (isinstance(iprot, THeaderProtocol.THeaderProtocolAccelerate) and iprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_BINARY_PROTOCOL)) and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None and fastproto is not None:
      fastproto.decode(self, iprot.trans, [self.__class__, self.thrift_spec, False], utf8strings=UTF8STRINGS, protoid=0)
      return
    if (isinstance(iprot, TCompactProtocol.TCompactProtocolAccelerated) or (isinstance(iprot, THeaderProtocol.THeaderProtocolAccelerate) and iprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_COMPACT_PROTOCOL)) and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None and fastproto is not None:
      fastproto.decode(self, iprot.trans, [self.__class__, self.thrift_spec, False], utf8strings=UTF8STRINGS, protoid=2)
      return
    iprot.readStructBegin()
    while True:
      (fname, ftype, fid) = iprot.readFieldBegin()
      if ftype == TType.STOP:
        break
      if fid == 1:
        if ftype == TType.I64:
          self.id = iprot.readI64()
        else:
          iprot.skip(ftype)
      elif fid == 2:
        if ftype == TType.STRING:
          self.password = iprot.readString().decode('utf-8') if UTF8STRINGS else iprot.readString()
        else:
          iprot.skip(ftype)
      else:
        iprot.skip(ftype)
      iprot.readFieldEnd()
    iprot.readStructEnd()

  def write(self, oprot):
    if (isinstance(oprot, TBinaryProtocol.TBinaryProtocolAccelerated) or (isinstance(oprot, THeaderProtocol.THeaderProtocolAccelerate) and oprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_BINARY_PROTOCOL)) and self.thrift_spec is not None and fastproto is not None:
      oprot.trans.write(fastproto.encode(self, [self.__class__, self.thrift_spec, False], utf8strings=UTF8STRINGS, protoid=0))
      return
    if (isinstance(oprot, TCompactProtocol.TCompactProtocolAccelerated) or (isinstance(oprot, THeaderProtocol.THeaderProtocolAccelerate) and oprot.get_protocol_id() == THeaderProtocol.THeaderProtocol.T_COMPACT_PROTOCOL)) and self.thrift_spec is not None and fastproto is not None:
      oprot.trans.write(fastproto.encode(self, [self.__class__, self.thrift_spec, False], utf8strings=UTF8STRINGS, protoid=2))
      return
    oprot.writeStructBegin('SecretStruct')
    if self.id != None:
      oprot.writeFieldBegin('id', TType.I64, 1)
      oprot.writeI64(self.id)
      oprot.writeFieldEnd()
    if self.password != None:
      oprot.writeFieldBegin('password', TType.STRING, 2)
      oprot.writeString(self.password.encode('utf-8')) if UTF8STRINGS and not isinstance(self.password, bytes) else oprot.writeString(self.password)
      oprot.writeFieldEnd()
    oprot.writeFieldStop()
    oprot.writeStructEnd()

  def readFromJson(self, json, is_text=True, **kwargs):
    kwargs_copy = dict(kwargs)
    relax_enum_validation = bool(kwargs_copy.pop('relax_enum_validation', False))
    set_cls = kwargs_copy.pop('custom_set_cls', set)
    dict_cls = kwargs_copy.pop('custom_dict_cls', dict)
    wrap_enum_constants = kwargs_copy.pop('wrap_enum_constants', False)
    if wrap_enum_constants and relax_enum_validation:
        raise ValueError(
            'wrap_enum_constants cannot be used together with relax_enum_validation'
        )
    if kwargs_copy:
        extra_kwargs = ', '.join(kwargs_copy.keys())
        raise ValueError(
            'Unexpected keyword arguments: ' + extra_kwargs
        )
    json_obj = json
    if is_text:
      json_obj = loads(json)
    if 'id' in json_obj and json_obj['id'] is not None:
      self.id = long(json_obj['id'])
    if 'password' in json_obj and json_obj['password'] is not None:
      self.password = json_obj['password']

  def __repr__(self):
    L = []
    padding = ' ' * 4
    if self.id is not None:
      value = pprint.pformat(self.id, indent=0)
      value = padding.join(value.splitlines(True))
      L.append('    id=%s' % (value))
    if self.password is not None:
      value = pprint.pformat(self.password, indent=0)
      value = padding.join(value.splitlines(True))
      L.append('    password=%s' % (value))
    return "%s(%s)" % (self.__class__.__name__, "\n" + ",\n".join(L) if L else '')

  def __eq__(self, other):
    if not isinstance(other, self.__class__):
      return False

    return self.__dict__ == other.__dict__ 

  def __ne__(self, other):
    return not (self == other)

  def __dir__(self):
    return (
      'id',
      'password',
    )

  __hash__ = object.__hash__

  def _to_python(self):
    import importlib
    import thrift.python.converter
    python_types = importlib.import_module("module.thrift_types")
    return thrift.python.converter.to_python_struct(python_types.SecretStruct, self)

  def _to_mutable_python(self):
    import importlib
    import thrift.python.mutable_converter
    python_mutable_types = importlib.import_module("module.thrift_mutable_types")
    return thrift.python.mutable_converter.to_mutable_python_struct_or_union(python_mutable_types.SecretStruct, self)

  def _to_py3(self):
    import importlib
    import thrift.py3.converter
    py3_types = importlib.import_module("module.types")
    return thrift.py3.converter.to_py3_struct(py3_types.SecretStruct, self)

  def _to_py_deprecated(self):
    return self

AwesomeStruct = MyStruct
FantasticStruct = MyStruct
list_string_6884 = UnimplementedTypedef()
all_structs.append(MyStructNestedAnnotation)
MyStructNestedAnnotation.thrift_spec = tuple(__EXPAND_THRIFT_SPEC((
  (1, TType.STRING, 'name', True, None, 2, ), # 1
)))

MyStructNestedAnnotation.thrift_struct_annotations = {
}
MyStructNestedAnnotation.thrift_field_annotations = {
}

def MyStructNestedAnnotation__init__(self, name=None,):
  self.name = name

MyStructNestedAnnotation.__init__ = MyStructNestedAnnotation__init__

def MyStructNestedAnnotation__setstate__(self, state):
  state.setdefault('name', None)
  self.__dict__ = state

MyStructNestedAnnotation.__getstate__ = lambda self: self.__dict__.copy()
MyStructNestedAnnotation.__setstate__ = MyStructNestedAnnotation__setstate__

all_structs.append(MyUnion)
MyUnion.thrift_spec = tuple(__EXPAND_THRIFT_SPEC((
)))

MyUnion.thrift_struct_annotations = {
}
MyUnion.thrift_field_annotations = {
}

all_structs.append(MyException)
MyException.thrift_spec = tuple(__EXPAND_THRIFT_SPEC((
)))

MyException.thrift_struct_annotations = {
}
MyException.thrift_field_annotations = {
}

all_structs.append(MyStruct)
MyStruct.thrift_spec = tuple(__EXPAND_THRIFT_SPEC((
  (1, TType.STRING, 'abstract', True, None, 2, ), # 1
  (2, TType.I64, 'major', None, None, 2, ), # 2
  (3, TType.STRING, 'annotation_with_quote', True, None, 2, ), # 3
  (4, TType.STRING, 'class_', True, None, 2, ), # 4
  (5, TType.STRING, 'annotation_with_trailing_comma', True, None, 2, ), # 5
  (6, TType.STRING, 'empty_annotations', True, None, 2, ), # 6
  (7, TType.I32, 'my_enum', MyEnum, None, 2, ), # 7
  (8, TType.LIST, 'cpp_type_annotation', (TType.STRING,True), None, 2, ), # 8
  (9, TType.STRUCT, 'my_union', [MyUnion, MyUnion.thrift_spec, True], None, 2, ), # 9
)))

MyStruct.thrift_struct_annotations = {
  "android.generate_builder": "1",
  "thrift.uri": "facebook.com/thrift/compiler/test/fixtures/basic-annotations/src/module/MyStruct",
}
MyStruct.thrift_field_annotations = {
  1: {
    "java.swift.name": "_abstract",
  },
  4: {
    "java.swift.name": "class_",
  },
  5: {
    "custom": "test",
  },
}

def MyStruct__init__(self, major=None, abstract=None, annotation_with_quote=None, class_=None, annotation_with_trailing_comma=None, empty_annotations=None, my_enum=None, cpp_type_annotation=None, my_union=None,):
  self.major = major
  self.abstract = abstract
  self.annotation_with_quote = annotation_with_quote
  self.class_ = class_
  self.annotation_with_trailing_comma = annotation_with_trailing_comma
  self.empty_annotations = empty_annotations
  self.my_enum = my_enum
  self.cpp_type_annotation = cpp_type_annotation
  self.my_union = my_union

MyStruct.__init__ = MyStruct__init__

def MyStruct__setstate__(self, state):
  state.setdefault('major', None)
  state.setdefault('abstract', None)
  state.setdefault('annotation_with_quote', None)
  state.setdefault('class_', None)
  state.setdefault('annotation_with_trailing_comma', None)
  state.setdefault('empty_annotations', None)
  state.setdefault('my_enum', None)
  state.setdefault('cpp_type_annotation', None)
  state.setdefault('my_union', None)
  self.__dict__ = state

MyStruct.__getstate__ = lambda self: self.__dict__.copy()
MyStruct.__setstate__ = MyStruct__setstate__

all_structs.append(SecretStruct)
SecretStruct.thrift_spec = tuple(__EXPAND_THRIFT_SPEC((
  (1, TType.I64, 'id', None, None, 2, ), # 1
  (2, TType.STRING, 'password', True, None, 2, ), # 2
)))

SecretStruct.thrift_struct_annotations = {
}
SecretStruct.thrift_field_annotations = {
  2: {
    "java.sensitive": "1",
  },
}

def SecretStruct__init__(self, id=None, password=None,):
  self.id = id
  self.password = password

SecretStruct.__init__ = SecretStruct__init__

def SecretStruct__setstate__(self, state):
  state.setdefault('id', None)
  state.setdefault('password', None)
  self.__dict__ = state

SecretStruct.__getstate__ = lambda self: self.__dict__.copy()
SecretStruct.__setstate__ = SecretStruct__setstate__

fix_spec(all_structs)
del all_structs
