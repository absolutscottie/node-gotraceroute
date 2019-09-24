#include <node.h>
#include <string>
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
    if (args.Length() < 4) {
      // Throw an Error that is passed back to JavaScript
      isolate->ThrowException(Exception::TypeError(
          String::NewFromUtf8(isolate, "Wrong number of arguments")));
      return;
    }

    // Check the argument types
    if (!args[1]->IsNumber() || !args[2]->IsNumber() || !args[3]->IsNumber()) {
      isolate->ThrowException(Exception::TypeError(
          String::NewFromUtf8(isolate, "Wrong arguments")));
      return;
    }

    String::Utf8Value hostname(args[0]->ToString());
    Local<String> result = String::NewFromUtf8(
                            isolate,
                            PingHost(
                              GoString{*hostname, args[0]->ToString()->Length()},
                              args[1]->NumberValue(),
                              args[2]->NumberValue(),
                              args[3]->NumberValue()));

    // Set the return value (using the passed in
    // FunctionCallbackInfo<Value>&)
    args.GetReturnValue().Set(result);
  }

  void init(Local<Object> exports) {
    NODE_SET_METHOD(exports, "pingHost", pingHost);
  }

  NODE_MODULE(pingModule, init)
}
