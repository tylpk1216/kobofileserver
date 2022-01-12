#include <QCoreApplication>
#include <QTextStream>
#include <QTimer>

#include <dlfcn.h>
#include <unistd.h>

#include "cli.h"

typedef void PlugWorkflowManager;

int Cli::immportBooks(void)
{
    QTextStream methodOut(stdout);

    void *fHandle;

    fHandle = dlopen("/usr/local/Kobo/libnickel.so.1.0.0", RTLD_LAZY);
    if (!fHandle) {
        methodOut << dlerror() << endl;
        return 1;
    }

    //libnickel 4.13.12638 * _ZN19PlugWorkflowManager14sharedInstanceEv
    PlugWorkflowManager *(*PlugWorkflowManager_sharedInstance)();
    PlugWorkflowManager_sharedInstance = (PlugWorkflowManager *(*)())dlsym(fHandle, "_ZN19PlugWorkflowManager14sharedInstanceEv");

    if (!PlugWorkflowManager_sharedInstance) {
        methodOut << dlerror() << endl;
        dlclose(fHandle);
        return 2;
    }

    PlugWorkflowManager *pWM = PlugWorkflowManager_sharedInstance();
    if (!pWM) {
        methodOut << "get PlugWorkflowManager pointer error\n" << endl;
        dlclose(fHandle);
        return 3;
    }

    /*
    //libnickel 4.13.12638 * _ZN19PlugWorkflowManager4syncEv
    void (*PlugWorkflowManager_sync)(PlugWorkflowManager*);
    PlugWorkflowManager_sync = (void (*)(PlugWorkflowManager*))dlsym(fHandle, "_ZN19PlugWorkflowManager4syncEv");

    if (!PlugWorkflowManager_sync) {
        methodOut << dlerror() << endl;
        dlclose(fHandle);
        return 4;
    }

    PlugWorkflowManager_sync(pWM);
    */

    //libnickel 4.13.12638 * _ZN19PlugWorkflowManager18onCancelAndConnectEv
    void (*PlugWorkflowManager_onCancelAndConnect)(PlugWorkflowManager*);
    PlugWorkflowManager_onCancelAndConnect = (void(*)(PlugWorkflowManager*))dlsym(fHandle, "_ZN19PlugWorkflowManager18onCancelAndConnectEv");

    if (!PlugWorkflowManager_onCancelAndConnect) {
        methodOut << dlerror() << endl;
        dlclose(fHandle);
        return 5;
    }

    //libnickel 4.13.12638 * _ZN19PlugWorkflowManager9unpluggedEv
    void (*PlugWorkflowManager_unplugged)(PlugWorkflowManager*);
    PlugWorkflowManager_unplugged = (void(*)(PlugWorkflowManager*))dlsym(fHandle, "_ZN19PlugWorkflowManager9unpluggedEv");

    if (!PlugWorkflowManager_unplugged) {
        methodOut << dlerror() << endl;
        dlclose(fHandle);
        return 6;
    }

    PlugWorkflowManager_onCancelAndConnect(pWM);
    sleep(1);
    PlugWorkflowManager_unplugged(pWM);

    dlclose(fHandle);
    return 0;
}

Cli::Cli(QObject* parent, int sec) : QObject(parent) {
    timeoutSec = sec;
}

void Cli::handleTimeout() {
    QCoreApplication::exit(99);
}

void Cli::start() {
    int res = immportBooks();
    if (res != 0) {
        QCoreApplication::exit(res);
    }

    if (timeoutSec > 0) {
        QTimer::singleShot(timeoutSec * 1000, this, SLOT(handleTimeout()));
    } else {
        QCoreApplication::quit();
    }
}