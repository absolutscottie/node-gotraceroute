cmd_Makefile := cd ..; /usr/local/lib/node_modules/node-gyp/gyp/gyp_main.py -fmake --ignore-environment "--toplevel-dir=." "-I/Users/scottie/Desktop/Node Module Test/src/build/config.gypi" -I/usr/local/lib/node_modules/node-gyp/addon.gypi -I/Users/scottie/.node-gyp/10.15.3/include/node/common.gypi "--depth=." "-Goutput_dir=." "--generator-output=build" "-Dlibrary=shared_library" "-Dvisibility=default" "-Dnode_root_dir=/Users/scottie/.node-gyp/10.15.3" "-Dnode_gyp_dir=/usr/local/lib/node_modules/node-gyp" "-Dnode_lib_file=/Users/scottie/.node-gyp/10.15.3/<(target_arch)/node.lib" "-Dmodule_root_dir=/Users/scottie/Desktop/Node Module Test/src" "-Dnode_engine=v8" binding.gyp
