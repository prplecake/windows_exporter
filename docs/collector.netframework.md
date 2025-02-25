# netframework collector

The netframework collector exposes metrics about dotnet framework.

|                         |                                                                                                                                                                                                                                                                                                                                                                                                        |
|-------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **Metric name prefix**  | `netframework_clrexceptions`                                                                                                                                                                                                                                                                                                                                                                           |
| **Classes**             | `Win32_PerfRawData_NETFramework_NETCLRExceptions`, `Win32_PerfRawData_NETFramework_NETCLRInterop`, `Win32_PerfRawData_NETFramework_NETCLRJit`, `Win32_PerfRawData_NETFramework_NETCLRLoading`, `Win32_PerfRawData_NETFramework_NETCLRLocksAndThreads`, `Win32_PerfRawData_NETFramework_NETCLRMemory`, `Win32_PerfRawData_NETFramework_NETCLRRemoting`, `Win32_PerfRawData_NETFramework_NETCLRSecurity` |
| **Enabled by default?** | No                                                                                                                                                                                                                                                                                                                                                                                                     |

## Flags

### `--collector.netframework.enabled`

Comma-separated list of collectors to use. Defaults to all, if not specified.

## Metrics

### CLR Exceptions

| Name                                                            | Description                                                                                                                                                                               | Type    | Labels    |
|-----------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|-----------|
| `windows_netframework_clrexceptions_exceptions_thrown_total`    | Displays the total number of exceptions thrown since the application started. This includes both .NET exceptions and unmanaged exceptions that are converted into .NET exceptions.        | counter | `process` |
| `windows_netframework_clrexceptions_exceptions_filters_total`   | Displays the total number of .NET exception filters executed. An exception filter evaluates regardless of whether an exception is handled.                                                | counter | `process` |
| `windows_netframework_clrexceptions_exceptions_finallys_total`  | Displays the total number of finally blocks executed. Only the finally blocks executed for an exception are counted; finally blocks on normal code paths are not counted by this counter. | counter | `process` |
| `windows_netframework_clrexceptions_throw_to_catch_depth_total` | Displays the total number of stack frames traversed, from the frame that threw the exception to the frame that handled the exception.                                                     | counter | `process` |

### CLR Interop

| Name                                                          | Description                                                                                                                                                                                                                                        | Type    | Labels    |
|---------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|-----------|
| `windows_netframework_clrinterop_com_callable_wrappers_total` | Displays the current number of COM callable wrappers (CCWs). A CCW is a proxy for a managed object being referenced from an unmanaged COM client.                                                                                                  | counter | `process` |
| `windows_netframework_clrinterop_interop_marshalling_total`   | Displays the total number of times arguments and return values have been marshaled from managed to unmanaged code, and vice versa, since the application started.                                                                                  | counter | `process` |
| `windows_netframework_clrinterop_interop_stubs_created_total` | Displays the current number of stubs created by the common language runtime. Stubs are responsible for marshaling arguments and return values from managed to unmanaged code, and vice versa, during a COM interop call or a platform invoke call. | counter | `process` |

### CLR JIT

| Name                                                      | Description                                                                                                                                                                                                           | Type    | Labels    |
|-----------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|-----------|
| `windows_netframework_clrjit_jit_methods_total`           | Displays the total number of methods JIT-compiled since the application started. This counter does not include pre-JIT-compiled methods.                                                                              | counter | `process` |
| `windows_netframework_clrjit_jit_time_percent`            | Displays the percentage of time spent in JIT compilation. This counter is updated at the end of every JIT compilation phase. A JIT compilation phase occurs when a method and its dependencies are compiled.          | gauge   | `process` |
| `windows_netframework_clrjit_jit_standard_failures_total` | Displays the peak number of methods the JIT compiler has failed to compile since the application started. This failure can occur if the MSIL cannot be verified or if there is an internal error in the JIT compiler. | counter | `process` |
| `windows_netframework_clrjit_jit_il_bytes_total`          | Displays the total number of Microsoft intermediate language (MSIL) bytes compiled by the just-in-time (JIT) compiler since the application started                                                                   | counter | `process` |

### CLR Loading

| Name                                                        | Description                                                                                                                                                                                                                                 | Type    | Labels    |
|-------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|-----------|
| `windows_netframework_clrloading_loader_heap_size_bytes`    | Displays the current size, in bytes, of the memory committed by the class loader across all application domains. Committed memory is the physical space reserved in the disk paging file.                                                   | gauge   | `process` |
| `windows_netframework_clrloading_appdomains_loaded_current` | Displays the current number of application domains loaded in this application.                                                                                                                                                              | gauge   | `process` |
| `windows_netframework_clrloading_assemblies_loaded_current` | Displays the current number of assemblies loaded across all application domains in the currently running application. If the assembly is loaded as domain-neutral from multiple application domains, this counter is incremented only once. | gauge   | `process` |
| `windows_netframework_clrloading_classes_loaded_current`    | Displays the current number of classes loaded in all assemblies.                                                                                                                                                                            | gauge   | `process` |
| `windows_netframework_clrloading_appdomains_loaded_total`   | Displays the peak number of application domains loaded since the application started.                                                                                                                                                       | counter | `process` |
| `windows_netframework_clrloading_appdomains_unloaded_total` | Displays the total number of application domains unloaded since the application started. If an application domain is loaded and unloaded multiple times, this counter increments each time the application domain is unloaded.              | counter | `process` |
| `windows_netframework_clrloading_assemblies_loaded_total`   | Displays the total number of assemblies loaded since the application started. If the assembly is loaded as domain-neutral from multiple application domains, this counter is incremented only once.                                         | counter | `process` |
| `windows_netframework_clrloading_classes_loaded_total`      | Displays the cumulative number of classes loaded in all assemblies since the application started.                                                                                                                                           | counter | `process` |
| `windows_netframework_clrloading_class_load_failures_total` | Displays the peak number of classes that have failed to load since the application started.                                                                                                                                                 | counter | `process` |

### CLR Locks and Threads

| Name                                                                 | Description                                                                                                                                                                                                                                                                                                                       | Type    | Labels    |
|----------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|-----------|
| `windows_netframework_clrlocksandthreads_current_queue_length`       | Displays the total number of threads that are currently waiting to acquire a managed lock in the application.                                                                                                                                                                                                                     | gauge   | `process` |
| `windows_netframework_clrlocksandthreads_current_logical_threads`    | Displays the number of current managed thread objects in the application. This counter maintains the count of both running and stopped threads.                                                                                                                                                                                   | gauge   | `process` |
| `windows_netframework_clrlocksandthreads_physical_threads_current`   | Displays the number of native operating system threads created and owned by the common language runtime to act as underlying threads for managed thread objects. This counter's value does not include the threads used by the runtime in its internal operations; it is a subset of the threads in the operating system process. | gauge   | `process` |
| `windows_netframework_clrlocksandthreads_recognized_threads_current` | Displays the number of threads that are currently recognized by the runtime. These threads are associated with a corresponding managed thread object. The runtime does not create these threads, but they have run inside the runtime at least once.                                                                              | gauge   | `process` |
| `windows_netframework_clrlocksandthreads_recognized_threads_total`   | Displays the total number of threads that have been recognized by the runtime since the application started. These threads are associated with a corresponding managed thread object. The runtime does not create these threads, but they have run inside the runtime at least once.                                              | counter | `process` |
| `windows_netframework_clrlocksandthreads_queue_length_total`         | Displays the total number of threads that waited to acquire a managed lock since the application started.                                                                                                                                                                                                                         | counter | `process` |
| `windows_netframework_clrlocksandthreads_contentions_total`          | Displays the total number of times that threads in the runtime have attempted to acquire a managed lock unsuccessfully.                                                                                                                                                                                                           | counter | `process` |

### CLR Memory

| Name                                                     | Description                                                                                                                                                                                                                                                       | Type    | Labels    |
|----------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|-----------|
| `windows_netframework_clrmemory_allocated_bytes_total`   | Displays the total number of bytes allocated on the garbage collection heap.                                                                                                                                                                                      | counter | `process` |
| `windows_netframework_clrmemory_finalization_survivors`  | Displays the number of garbage-collected objects that survive a collection because they are waiting to be finalized.                                                                                                                                              | gauge   | `process` |
| `windows_netframework_clrmemory_heap_size_bytes`         | Displays the maximum bytes that can be allocated; it does not indicate the current number of bytes allocated.                                                                                                                                                     | gauge   | `process` |
| `windows_netframework_clrmemory_promoted_bytes`          | Displays the bytes that were promoted from the generation to the next one during the last GC. Memory is promoted when it survives a garbage collection.                                                                                                           | gauge   | `process` |
| `windows_netframework_clrmemory_number_gc_handles`       | Displays the current number of garbage collection handles in use. Garbage collection handles are handles to resources external to the common language runtime and the managed environment.                                                                        | gauge   | `process` |
| `windows_netframework_clrmemory_collections_total`       | Displays the number of times the generation objects are garbage collected since the application started.                                                                                                                                                          | counter | `process` |
| `windows_netframework_clrmemory_induced_gc_total`        | Displays the peak number of times garbage collection was performed because of an explicit call to GC.Collect.                                                                                                                                                     | counter | `process` |
| `windows_netframework_clrmemory_number_pinned_objects`   | Displays the number of pinned objects encountered in the last garbage collection.                                                                                                                                                                                 | gauge   | `process` |
| `windows_netframework_clrmemory_number_sink_blocksinuse` | Displays the current number of synchronization blocks in use. Synchronization blocks are per-object data structures allocated for storing synchronization information. They hold weak references to managed objects and must be scanned by the garbage collector. | gauge   | `process` |
| `windows_netframework_clrmemory_committed_bytes`         | Displays the amount of virtual memory, in bytes, currently committed by the garbage collector. Committed memory is the physical memory for which space has been reserved in the disk paging file.                                                                 | gauge   | `process` |
| `windows_netframework_clrmemory_reserved_bytes`          | Displays the amount of virtual memory, in bytes, currently reserved by the garbage collector. Reserved memory is the virtual memory space reserved for the application when no disk or main memory pages have been used.                                          | gauge   | `process` |
| `windows_netframework_clrmemory_gc_time_percent`         | Displays the percentage of time that was spent performing a garbage collection in the last sample.                                                                                                                                                                | gauge   | `process` |

### CLR Remoting

| Name                                                            | Description                                                                                                         | Type    | Labels    |
|-----------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------|---------|-----------|
| `windows_netframework_clrremoting_channels_total`               | Displays the total number of remoting channels registered across all application domains since application started. | counter | `process` |
| `windows_netframework_clrremoting_context_bound_classes_loaded` | Displays the current number of context-bound classes that are loaded.                                               | gauge   | `process` |
| `windows_netframework_clrremoting_context_bound_objects_total`  | Displays the total number of context-bound objects allocated.                                                       | counter | `process` |
| `windows_netframework_clrremoting_context_proxies_total`        | Displays the total number of remoting proxy objects in this process since it started.                               | counter | `process` |
| `windows_netframework_clrremoting_contexts`                     | Displays the current number of remoting contexts in the application.                                                | gauge   | `process` |
| `windows_netframework_clrremoting_remote_calls_total`           | Displays the total number of remote procedure calls invoked since the application started.                          | counter | `process` |

### CLR Security

| Name                                                      | Description                                                                                               | Type    | Labels    |
|-----------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|---------|-----------|
| `windows_netframework_clrsecurity_link_time_checks_total` | Displays the total number of link-time code access security checks since the application started.         | counter | `process` |
| `windows_netframework_clrsecurity_rt_checks_time_percent` | Displays the percentage of time spent performing runtime code access security checks in the last sample.  | gauge   | `process` |
| `windows_netframework_clrsecurity_stack_walk_depth`       | Displays the depth of the stack during that last runtime code access security check.                      | gauge   | `process` |
| `windows_netframework_clrsecurity_runtime_checks_total`   | Displays the total number of runtime code access security checks performed since the application started. | counter | `process` |

### Example metric
_This collector does not yet have explained examples, we would appreciate your help adding them!_

## Useful queries
_This collector does not yet have any useful queries added, we would appreciate your help adding them!_

## Alerting examples
_This collector does not yet have alerting examples, we would appreciate your help adding them!_
