/**
 * @fileoverview gRPC-Web generated client stub for api.v1
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.1
// 	protoc              v3.21.12
// source: src/chat.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')
const proto = {};
proto.api = {};
proto.api.v1 = require('./chat_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.api.v1.ChatClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.api.v1.ChatPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.v1.UserRegister,
 *   !proto.api.v1.UserRegisterReply>}
 */
const methodDescriptor_Chat_Register = new grpc.web.MethodDescriptor(
  '/api.v1.Chat/Register',
  grpc.web.MethodType.UNARY,
  proto.api.v1.UserRegister,
  proto.api.v1.UserRegisterReply,
  /**
   * @param {!proto.api.v1.UserRegister} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.v1.UserRegisterReply.deserializeBinary
);


/**
 * @param {!proto.api.v1.UserRegister} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.api.v1.UserRegisterReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.v1.UserRegisterReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.v1.ChatClient.prototype.register =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.v1.Chat/Register',
      request,
      metadata || {},
      methodDescriptor_Chat_Register,
      callback);
};


/**
 * @param {!proto.api.v1.UserRegister} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.v1.UserRegisterReply>}
 *     Promise that resolves to the response
 */
proto.api.v1.ChatPromiseClient.prototype.register =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.v1.Chat/Register',
      request,
      metadata || {},
      methodDescriptor_Chat_Register);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.v1.Null,
 *   !proto.api.v1.UserListReply>}
 */
const methodDescriptor_Chat_ListUsers = new grpc.web.MethodDescriptor(
  '/api.v1.Chat/ListUsers',
  grpc.web.MethodType.UNARY,
  proto.api.v1.Null,
  proto.api.v1.UserListReply,
  /**
   * @param {!proto.api.v1.Null} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.v1.UserListReply.deserializeBinary
);


/**
 * @param {!proto.api.v1.Null} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.api.v1.UserListReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.v1.UserListReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.v1.ChatClient.prototype.listUsers =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.v1.Chat/ListUsers',
      request,
      metadata || {},
      methodDescriptor_Chat_ListUsers,
      callback);
};


/**
 * @param {!proto.api.v1.Null} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.v1.UserListReply>}
 *     Promise that resolves to the response
 */
proto.api.v1.ChatPromiseClient.prototype.listUsers =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.v1.Chat/ListUsers',
      request,
      metadata || {},
      methodDescriptor_Chat_ListUsers);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.v1.NewMessage,
 *   !proto.api.v1.Null>}
 */
const methodDescriptor_Chat_Message = new grpc.web.MethodDescriptor(
  '/api.v1.Chat/Message',
  grpc.web.MethodType.UNARY,
  proto.api.v1.NewMessage,
  proto.api.v1.Null,
  /**
   * @param {!proto.api.v1.NewMessage} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.v1.Null.deserializeBinary
);


/**
 * @param {!proto.api.v1.NewMessage} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.api.v1.Null)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.v1.Null>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.v1.ChatClient.prototype.message =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.v1.Chat/Message',
      request,
      metadata || {},
      methodDescriptor_Chat_Message,
      callback);
};


/**
 * @param {!proto.api.v1.NewMessage} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.v1.Null>}
 *     Promise that resolves to the response
 */
proto.api.v1.ChatPromiseClient.prototype.message =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.v1.Chat/Message',
      request,
      metadata || {},
      methodDescriptor_Chat_Message);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.v1.Null,
 *   !proto.api.v1.PollReply>}
 */
const methodDescriptor_Chat_Poll = new grpc.web.MethodDescriptor(
  '/api.v1.Chat/Poll',
  grpc.web.MethodType.UNARY,
  proto.api.v1.Null,
  proto.api.v1.PollReply,
  /**
   * @param {!proto.api.v1.Null} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.v1.PollReply.deserializeBinary
);


/**
 * @param {!proto.api.v1.Null} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.api.v1.PollReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.v1.PollReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.v1.ChatClient.prototype.poll =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.v1.Chat/Poll',
      request,
      metadata || {},
      methodDescriptor_Chat_Poll,
      callback);
};


/**
 * @param {!proto.api.v1.Null} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.v1.PollReply>}
 *     Promise that resolves to the response
 */
proto.api.v1.ChatPromiseClient.prototype.poll =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.v1.Chat/Poll',
      request,
      metadata || {},
      methodDescriptor_Chat_Poll);
};


module.exports = proto.api.v1;
