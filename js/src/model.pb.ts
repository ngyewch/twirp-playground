// THIS IS AN AUTOGENERATED FILE. DO NOT EDIT THIS FILE DIRECTLY.
// Source: model.proto
/* eslint-disable */

import type { ByteSource, PartialDeep } from "protoscript";
import * as protoscript from "protoscript";

//========================================//
//                 Types                  //
//========================================//

export interface ToUpperRequest {
  text: string;
}

export interface ToUpperResponse {
  text: string;
}

//========================================//
//        Protobuf Encode / Decode        //
//========================================//

export const ToUpperRequest = {
  /**
   * Serializes ToUpperRequest to protobuf.
   */
  encode: function (msg: PartialDeep<ToUpperRequest>): Uint8Array {
    return ToUpperRequest._writeMessage(
      msg,
      new protoscript.BinaryWriter(),
    ).getResultBuffer();
  },

  /**
   * Deserializes ToUpperRequest from protobuf.
   */
  decode: function (bytes: ByteSource): ToUpperRequest {
    return ToUpperRequest._readMessage(
      ToUpperRequest.initialize(),
      new protoscript.BinaryReader(bytes),
    );
  },

  /**
   * Initializes ToUpperRequest with all fields set to their default value.
   */
  initialize: function (msg?: Partial<ToUpperRequest>): ToUpperRequest {
    return {
      text: "",
      ...msg,
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: PartialDeep<ToUpperRequest>,
    writer: protoscript.BinaryWriter,
  ): protoscript.BinaryWriter {
    if (msg.text) {
      writer.writeString(1, msg.text);
    }
    return writer;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: ToUpperRequest,
    reader: protoscript.BinaryReader,
  ): ToUpperRequest {
    while (reader.nextField()) {
      const field = reader.getFieldNumber();
      switch (field) {
        case 1: {
          msg.text = reader.readString();
          break;
        }
        default: {
          reader.skipField();
          break;
        }
      }
    }
    return msg;
  },
};

export const ToUpperResponse = {
  /**
   * Serializes ToUpperResponse to protobuf.
   */
  encode: function (msg: PartialDeep<ToUpperResponse>): Uint8Array {
    return ToUpperResponse._writeMessage(
      msg,
      new protoscript.BinaryWriter(),
    ).getResultBuffer();
  },

  /**
   * Deserializes ToUpperResponse from protobuf.
   */
  decode: function (bytes: ByteSource): ToUpperResponse {
    return ToUpperResponse._readMessage(
      ToUpperResponse.initialize(),
      new protoscript.BinaryReader(bytes),
    );
  },

  /**
   * Initializes ToUpperResponse with all fields set to their default value.
   */
  initialize: function (msg?: Partial<ToUpperResponse>): ToUpperResponse {
    return {
      text: "",
      ...msg,
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: PartialDeep<ToUpperResponse>,
    writer: protoscript.BinaryWriter,
  ): protoscript.BinaryWriter {
    if (msg.text) {
      writer.writeString(1, msg.text);
    }
    return writer;
  },

  /**
   * @private
   */
  _readMessage: function (
    msg: ToUpperResponse,
    reader: protoscript.BinaryReader,
  ): ToUpperResponse {
    while (reader.nextField()) {
      const field = reader.getFieldNumber();
      switch (field) {
        case 1: {
          msg.text = reader.readString();
          break;
        }
        default: {
          reader.skipField();
          break;
        }
      }
    }
    return msg;
  },
};

//========================================//
//          JSON Encode / Decode          //
//========================================//

export const ToUpperRequestJSON = {
  /**
   * Serializes ToUpperRequest to JSON.
   */
  encode: function (msg: PartialDeep<ToUpperRequest>): string {
    return JSON.stringify(ToUpperRequestJSON._writeMessage(msg));
  },

  /**
   * Deserializes ToUpperRequest from JSON.
   */
  decode: function (json: string): ToUpperRequest {
    return ToUpperRequestJSON._readMessage(
      ToUpperRequestJSON.initialize(),
      JSON.parse(json),
    );
  },

  /**
   * Initializes ToUpperRequest with all fields set to their default value.
   */
  initialize: function (msg?: Partial<ToUpperRequest>): ToUpperRequest {
    return {
      text: "",
      ...msg,
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: PartialDeep<ToUpperRequest>,
  ): Record<string, unknown> {
    const json: Record<string, unknown> = {};
    if (msg.text) {
      json["text"] = msg.text;
    }
    return json;
  },

  /**
   * @private
   */
  _readMessage: function (msg: ToUpperRequest, json: any): ToUpperRequest {
    const _text_ = json["text"];
    if (_text_) {
      msg.text = _text_;
    }
    return msg;
  },
};

export const ToUpperResponseJSON = {
  /**
   * Serializes ToUpperResponse to JSON.
   */
  encode: function (msg: PartialDeep<ToUpperResponse>): string {
    return JSON.stringify(ToUpperResponseJSON._writeMessage(msg));
  },

  /**
   * Deserializes ToUpperResponse from JSON.
   */
  decode: function (json: string): ToUpperResponse {
    return ToUpperResponseJSON._readMessage(
      ToUpperResponseJSON.initialize(),
      JSON.parse(json),
    );
  },

  /**
   * Initializes ToUpperResponse with all fields set to their default value.
   */
  initialize: function (msg?: Partial<ToUpperResponse>): ToUpperResponse {
    return {
      text: "",
      ...msg,
    };
  },

  /**
   * @private
   */
  _writeMessage: function (
    msg: PartialDeep<ToUpperResponse>,
  ): Record<string, unknown> {
    const json: Record<string, unknown> = {};
    if (msg.text) {
      json["text"] = msg.text;
    }
    return json;
  },

  /**
   * @private
   */
  _readMessage: function (msg: ToUpperResponse, json: any): ToUpperResponse {
    const _text_ = json["text"];
    if (_text_) {
      msg.text = _text_;
    }
    return msg;
  },
};