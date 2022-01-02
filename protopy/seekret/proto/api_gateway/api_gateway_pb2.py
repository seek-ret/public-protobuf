# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: seekret/proto/api_gateway/api_gateway.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='seekret/proto/api_gateway/api_gateway.proto',
  package='seekret.api_gateway',
  syntax='proto3',
  serialized_options=b'Z7github.com/seek-ret/public-protobuf/protogo/api_gateway',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n+seekret/proto/api_gateway/api_gateway.proto\x12\x13seekret.api_gateway\"\x9b\x01\n\x0eWorkspaceEntry\x12\x10\n\x08\x62\x61se_dir\x18\x01 \x01(\t\x12\x13\n\x0b\x62ucket_name\x18\x02 \x01(\t\x12\x10\n\x08\x65ndpoint\x18\x03 \x01(\t\x12\x12\n\naccess_key\x18\x04 \x01(\t\x12\x12\n\nsecret_key\x18\x05 \x01(\t\x12\x10\n\x08insecure\x18\x06 \x01(\x08\x12\x16\n\x0ehost_whitelist\x18\x07 \x03(\t\"7\n\"FetchWorkspaceConfigurationRequest\x12\x11\n\tworkspace\x18\x01 \x01(\t\"Y\n#FetchWorkspaceConfigurationResponse\x12\x32\n\x05\x65ntry\x18\x01 \x01(\x0b\x32#.seekret.api_gateway.WorkspaceEntry2\x9f\x01\n\nApiGateway\x12\x90\x01\n\x1b\x46\x65tchWorkspaceConfiguration\x12\x37.seekret.api_gateway.FetchWorkspaceConfigurationRequest\x1a\x38.seekret.api_gateway.FetchWorkspaceConfigurationResponseB9Z7github.com/seek-ret/public-protobuf/protogo/api_gatewayb\x06proto3'
)




_WORKSPACEENTRY = _descriptor.Descriptor(
  name='WorkspaceEntry',
  full_name='seekret.api_gateway.WorkspaceEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='base_dir', full_name='seekret.api_gateway.WorkspaceEntry.base_dir', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='bucket_name', full_name='seekret.api_gateway.WorkspaceEntry.bucket_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='endpoint', full_name='seekret.api_gateway.WorkspaceEntry.endpoint', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='access_key', full_name='seekret.api_gateway.WorkspaceEntry.access_key', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='secret_key', full_name='seekret.api_gateway.WorkspaceEntry.secret_key', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='insecure', full_name='seekret.api_gateway.WorkspaceEntry.insecure', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='host_whitelist', full_name='seekret.api_gateway.WorkspaceEntry.host_whitelist', index=6,
      number=7, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=69,
  serialized_end=224,
)


_FETCHWORKSPACECONFIGURATIONREQUEST = _descriptor.Descriptor(
  name='FetchWorkspaceConfigurationRequest',
  full_name='seekret.api_gateway.FetchWorkspaceConfigurationRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='workspace', full_name='seekret.api_gateway.FetchWorkspaceConfigurationRequest.workspace', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=226,
  serialized_end=281,
)


_FETCHWORKSPACECONFIGURATIONRESPONSE = _descriptor.Descriptor(
  name='FetchWorkspaceConfigurationResponse',
  full_name='seekret.api_gateway.FetchWorkspaceConfigurationResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='entry', full_name='seekret.api_gateway.FetchWorkspaceConfigurationResponse.entry', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=283,
  serialized_end=372,
)

_FETCHWORKSPACECONFIGURATIONRESPONSE.fields_by_name['entry'].message_type = _WORKSPACEENTRY
DESCRIPTOR.message_types_by_name['WorkspaceEntry'] = _WORKSPACEENTRY
DESCRIPTOR.message_types_by_name['FetchWorkspaceConfigurationRequest'] = _FETCHWORKSPACECONFIGURATIONREQUEST
DESCRIPTOR.message_types_by_name['FetchWorkspaceConfigurationResponse'] = _FETCHWORKSPACECONFIGURATIONRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

WorkspaceEntry = _reflection.GeneratedProtocolMessageType('WorkspaceEntry', (_message.Message,), {
  'DESCRIPTOR' : _WORKSPACEENTRY,
  '__module__' : 'seekret.proto.api_gateway.api_gateway_pb2'
  # @@protoc_insertion_point(class_scope:seekret.api_gateway.WorkspaceEntry)
  })
_sym_db.RegisterMessage(WorkspaceEntry)

FetchWorkspaceConfigurationRequest = _reflection.GeneratedProtocolMessageType('FetchWorkspaceConfigurationRequest', (_message.Message,), {
  'DESCRIPTOR' : _FETCHWORKSPACECONFIGURATIONREQUEST,
  '__module__' : 'seekret.proto.api_gateway.api_gateway_pb2'
  # @@protoc_insertion_point(class_scope:seekret.api_gateway.FetchWorkspaceConfigurationRequest)
  })
_sym_db.RegisterMessage(FetchWorkspaceConfigurationRequest)

FetchWorkspaceConfigurationResponse = _reflection.GeneratedProtocolMessageType('FetchWorkspaceConfigurationResponse', (_message.Message,), {
  'DESCRIPTOR' : _FETCHWORKSPACECONFIGURATIONRESPONSE,
  '__module__' : 'seekret.proto.api_gateway.api_gateway_pb2'
  # @@protoc_insertion_point(class_scope:seekret.api_gateway.FetchWorkspaceConfigurationResponse)
  })
_sym_db.RegisterMessage(FetchWorkspaceConfigurationResponse)


DESCRIPTOR._options = None

_APIGATEWAY = _descriptor.ServiceDescriptor(
  name='ApiGateway',
  full_name='seekret.api_gateway.ApiGateway',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=375,
  serialized_end=534,
  methods=[
  _descriptor.MethodDescriptor(
    name='FetchWorkspaceConfiguration',
    full_name='seekret.api_gateway.ApiGateway.FetchWorkspaceConfiguration',
    index=0,
    containing_service=None,
    input_type=_FETCHWORKSPACECONFIGURATIONREQUEST,
    output_type=_FETCHWORKSPACECONFIGURATIONRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_APIGATEWAY)

DESCRIPTOR.services_by_name['ApiGateway'] = _APIGATEWAY

# @@protoc_insertion_point(module_scope)
