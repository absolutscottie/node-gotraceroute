#include <node.h>
#include "module.h"

namespace pingModule {

  using v8::FunctionCallbackInfo;
  using v8::Isolate;
  using v8::Local;
  using v8::Object;
  using v8::String;
  using v8::Value;
  using v8::Number;
  using v8::Exception;

  void pingHost(const FunctionCallbackInfo<Value>& args) {
    Isolate* isolate = args.GetIsolate();

    // Check the number of arguments passed.
    if (args.Length() < 2) {
      // Throw an Error that is passed back to JavaScript
      isolate->ThrowException(Exception::TypeError(
          String::NewFromUtf8(isolate, "Wrong number of arguments")));
      return;
    }

    // Check the argument types
    if (!args[1]->IsNumber()) {
      isolate->ThrowException(Exception::TypeError(
          String::NewFromUtf8(isolate, "Wrong arguments")));
      return;
    }

    v8::String::Utf8Value hostname(args[0]->ToString());

    // Perform the operation
    Local<Number> rtt = Number::New(
                          isolate,
                          PingHost(GoString{*hostname, args[0]->ToString()->Length()}, args[1]->NumberValue()));

    // Set the return value (using the passed in
    // FunctionCallbackInfo<Value>&)
    args.GetReturnValue().Set(rtt);
  }

  void init(Local<Object> exports) {
    NODE_SET_METHOD(exports, "pingHost", pingHost);
  }

  NODE_MODULE(pingModule, init)
}
