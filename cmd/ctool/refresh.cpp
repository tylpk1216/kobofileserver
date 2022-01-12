#include <QCoreApplication>
#include <QTimer>

#include <cstdlib>

#include "cli.h"

int main(int argc, char *argv[])
{
    QCoreApplication app(argc, argv);

    int timeoutSec = 0;
    if (argc == 2) {
        timeoutSec = atoi(argv[1]);
    }

    Cli cli(&app, timeoutSec);
    QTimer::singleShot(0, &cli, SLOT(start()));

    return app.exec();
}

