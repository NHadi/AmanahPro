const { app, BrowserWindow } = require('electron');

let mainWindow;

app.on('ready', () => {
    mainWindow = new BrowserWindow({
        webPreferences: {
            nodeIntegration: false // Secure browsing
        }
    });

    // Load the local web server URL
    mainWindow.loadURL('https://amanahpro.pilarasamandiri.com/project.html');
    // Maximize the window instead of fullscreen
    mainWindow.maximize();
    mainWindow.on('closed', () => {
        mainWindow = null;
    });
});

app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
        app.quit();
    }
});
