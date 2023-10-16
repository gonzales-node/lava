/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
var gogoproto_gogo_pb = require('../../../../gogoproto/gogo_pb.js');
goog.exportSymbol('proto.lavanet.lava.downtime.v1.Downtime', null, global);
goog.exportSymbol('proto.lavanet.lava.downtime.v1.Params', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.lavanet.lava.downtime.v1.Params = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.lavanet.lava.downtime.v1.Params, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.lavanet.lava.downtime.v1.Params.displayName = 'proto.lavanet.lava.downtime.v1.Params';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.lavanet.lava.downtime.v1.Params.prototype.toObject = function(opt_includeInstance) {
  return proto.lavanet.lava.downtime.v1.Params.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.lavanet.lava.downtime.v1.Params} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.downtime.v1.Params.toObject = function(includeInstance, msg) {
  var f, obj = {
    downtimeDuration: (f = msg.getDowntimeDuration()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
    epochDuration: (f = msg.getEpochDuration()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.lavanet.lava.downtime.v1.Params}
 */
proto.lavanet.lava.downtime.v1.Params.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.lavanet.lava.downtime.v1.Params;
  return proto.lavanet.lava.downtime.v1.Params.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.lavanet.lava.downtime.v1.Params} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.lavanet.lava.downtime.v1.Params}
 */
proto.lavanet.lava.downtime.v1.Params.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setDowntimeDuration(value);
      break;
    case 2:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setEpochDuration(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.lavanet.lava.downtime.v1.Params.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.lavanet.lava.downtime.v1.Params.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.lavanet.lava.downtime.v1.Params} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.downtime.v1.Params.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDowntimeDuration();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getEpochDuration();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
};


/**
 * optional google.protobuf.Duration downtime_duration = 1;
 * @return {?proto.google.protobuf.Duration}
 */
proto.lavanet.lava.downtime.v1.Params.prototype.getDowntimeDuration = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 1));
};


/** @param {?proto.google.protobuf.Duration|undefined} value */
proto.lavanet.lava.downtime.v1.Params.prototype.setDowntimeDuration = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.lavanet.lava.downtime.v1.Params.prototype.clearDowntimeDuration = function() {
  this.setDowntimeDuration(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.lavanet.lava.downtime.v1.Params.prototype.hasDowntimeDuration = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional google.protobuf.Duration epoch_duration = 2;
 * @return {?proto.google.protobuf.Duration}
 */
proto.lavanet.lava.downtime.v1.Params.prototype.getEpochDuration = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 2));
};


/** @param {?proto.google.protobuf.Duration|undefined} value */
proto.lavanet.lava.downtime.v1.Params.prototype.setEpochDuration = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.lavanet.lava.downtime.v1.Params.prototype.clearEpochDuration = function() {
  this.setEpochDuration(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.lavanet.lava.downtime.v1.Params.prototype.hasEpochDuration = function() {
  return jspb.Message.getField(this, 2) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.lavanet.lava.downtime.v1.Downtime = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.lavanet.lava.downtime.v1.Downtime, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.lavanet.lava.downtime.v1.Downtime.displayName = 'proto.lavanet.lava.downtime.v1.Downtime';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.lavanet.lava.downtime.v1.Downtime.prototype.toObject = function(opt_includeInstance) {
  return proto.lavanet.lava.downtime.v1.Downtime.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.lavanet.lava.downtime.v1.Downtime} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.downtime.v1.Downtime.toObject = function(includeInstance, msg) {
  var f, obj = {
    block: jspb.Message.getFieldWithDefault(msg, 1, 0),
    duration: (f = msg.getDuration()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.lavanet.lava.downtime.v1.Downtime}
 */
proto.lavanet.lava.downtime.v1.Downtime.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.lavanet.lava.downtime.v1.Downtime;
  return proto.lavanet.lava.downtime.v1.Downtime.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.lavanet.lava.downtime.v1.Downtime} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.lavanet.lava.downtime.v1.Downtime}
 */
proto.lavanet.lava.downtime.v1.Downtime.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setBlock(value);
      break;
    case 2:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setDuration(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.lavanet.lava.downtime.v1.Downtime.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.lavanet.lava.downtime.v1.Downtime.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.lavanet.lava.downtime.v1.Downtime} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.lavanet.lava.downtime.v1.Downtime.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getBlock();
  if (f !== 0) {
    writer.writeUint64(
      1,
      f
    );
  }
  f = message.getDuration();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
};


/**
 * optional uint64 block = 1;
 * @return {number}
 */
proto.lavanet.lava.downtime.v1.Downtime.prototype.getBlock = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.lavanet.lava.downtime.v1.Downtime.prototype.setBlock = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional google.protobuf.Duration duration = 2;
 * @return {?proto.google.protobuf.Duration}
 */
proto.lavanet.lava.downtime.v1.Downtime.prototype.getDuration = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 2));
};


/** @param {?proto.google.protobuf.Duration|undefined} value */
proto.lavanet.lava.downtime.v1.Downtime.prototype.setDuration = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.lavanet.lava.downtime.v1.Downtime.prototype.clearDuration = function() {
  this.setDuration(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.lavanet.lava.downtime.v1.Downtime.prototype.hasDuration = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto.lavanet.lava.downtime.v1);
