#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>

typedef void PlugWorkflowManager;

int main()
{
    void *fHandle;

    fHandle = dlopen("/user/local/Kobo/libnickel.so.1.0.0", RTLD_LAZY);
    if (!fHandle) {
        fprintf(stderr, "%s\n", dlerror());
        exit(1);
    }

    //libnickel 4.13.12638 * _ZN19PlugWorkflowManager14sharedInstanceEv
    PlugWorkflowManager *(*PlugWorkflowManager_sharedInstance)();
    PlugWorkflowManager_sharedInstance = (PlugWorkflowManager *(*)())dlsym(fHandle, "_ZN19PlugWorkflowManager14sharedInstanceEv");

    if (!PlugWorkflowManager_sharedInstance) {
        fprintf(stderr, "%s\n", dlerror());
        exit(1);
    }

    PlugWorkflowManager *pWM = PlugWorkflowManager_sharedInstance();
    if (!pWM) {
        fprintf(stderr, "get PlugWorkflowManager pointer error\n");
        exit(1);
    }

    //libnickel 4.13.12638 * _ZN19PlugWorkflowManager4syncEv
    void (*PlugWorkflowManager_sync)(PlugWorkflowManager*);
    PlugWorkflowManager_sync = (void (*)(PlugWorkflowManager*))dlsym(fHandle, " _ZN19PlugWorkflowManager4syncEv");

    if (!PlugWorkflowManager_sync) {
        fprintf(stderr, "%s\n", dlerror());
        exit(1);
    }

    PlugWorkflowManager_sync(pWM);

    dlclose(fHandle);
    return 0;
}



