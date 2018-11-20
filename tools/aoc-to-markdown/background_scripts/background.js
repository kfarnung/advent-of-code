browser.runtime.onMessage.addListener((message) => {
    if (message.action === "saveAs") {
        const blob = new Blob([message.text], { type: "text/plain;charset=utf-8" });

        browser.downloads.download({
            url: URL.createObjectURL(blob),
            filename: "README.md",
            saveAs: true,
        });
    }
});

browser.browserAction.onClicked.addListener(() => {
    console.log('Clicked');
    browser.tabs.executeScript({file: "/content_scripts/index.js"})
        .catch((e) => {
            console.error(e);
        });
});
