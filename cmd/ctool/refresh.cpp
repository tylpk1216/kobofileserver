#include <QCoreApplication>
#include <QTimer>

#include "cli.h"

int main(int argc, char *argv[])
{
    QCoreApplication app(argc, argv);

    Cli cli(&app);
    QTimer::singleShot(0, &cli, SLOT(start()));

    return app.exec();
}

