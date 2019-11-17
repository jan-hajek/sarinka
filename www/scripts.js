let port = location.port;

let url = location.protocol + "//" + location.hostname;
if (port !== "80") {
    url += ":" + port
}
let urlParams = new URLSearchParams(location.search);
let position = 0;
let channelId = urlParams.get("channelId");
let currentId = urlParams.get("id");
let nextId = null;
let player;
let screenWidth = screen.width;

function getChannelsUrl() {
    return url + "/channels/";
}


function startHomepage() {
    getJSON(getChannelsUrl(), function (status, response) {
        let channelsDiv = document.getElementById("channels");
        channelsDiv.innerHTML = "";

        response.Channels.map(function (value, index, array) {
            let a = document.createElement("a");
            a.classList.add("channel");
            a.href = "play/?channelId=" + value.Id;

            let img = new Image();
            img.src = value.ThumbnailUrl;
            img.classList.add("img");

            var text = document.createElement("div");
            text.innerHTML = value.Name;
            text.classList.add("title");

            a.appendChild(img);
            channelsDiv.appendChild(a);
        });
    });
}


function getPlayUrl(id) {
    return createUrl("/current/", id, channelId);
}

function getPreviewUrl(id) {
    return createUrl("/preview/", id, channelId);
}

function createUrl(prefix, id, channelId) {
    let data = {};

    let u = url + prefix;
    if (id !== null) {
        data.id = id
    }
    if (channelId !== null) {
        data.channelId = channelId
    }

    let searchParams = new URLSearchParams(data);

    return u + "?" + searchParams.toString()
}

function onYouTubeIframeAPIReady() {
    registerPlayOnKeydown();
    registerNightModeIcon();

    getJSON(getPlayUrl(currentId), function (status, response) {
        changePage(response)
    });
}

function nextVideo() {
    getJSON(getPlayUrl(nextId), function (status, response) {
        changePage(response)
        history.pushState({}, "xx", createUrl("/play/", currentId, channelId));
    })
}

function changePage(response) {
    if (response.Current != null) {
        currentId = response.Current.Id;
        document.title = response.Current.Title;
        setYoutubeId(currentId);
    }
    if (response.Next != null) {
        nextId = response.Next.Id;
        loadPreview(nextId);
    }
}

function setYoutubeId(id) {
    if (!player) {
        let width = screenWidth - 280;
        let height = screen.height;

        let overlay = document.getElementById("playerOverlay");
        overlay.style.width = width + "px";
        overlay.style.height = height + "px";

        player = new YT.Player('player', {
            width: width,
            height: height,
            videoId: id,
            playerVars: {start: 1, rel: 0, showinfo: 0, disablekb: 0, modestbranding: 1},
            events: {
                'onReady': function onPlayerReady(event) {
                    event.target.playVideo();
                },
                'onStateChange': onPlayerStateChange
            }
        });
    } else {
        player.loadVideoById(id, 1);
    }
}

let playing = false;

function onPlayerStateChange(event) {
    if (event.data === YT.PlayerState.ENDED) {
        nextVideo();
    }
    playing = event.data === YT.PlayerState.PLAYING;
}

function registerPlayOnKeydown() {
    document.onkeydown = function (event) {
        switch (event.key) {
            case "Enter":
                nextVideo();
                break;
            case " ":
                event.preventDefault();
                if (playing) {
                    player.pauseVideo();
                } else {
                    player.playVideo();
                }
                break;
            case "Escape":
                console.log("escape");
                window.location.href = url;
                return false;
        }
    };
}

function loadPreview(id) {
    getJSON(getPreviewUrl(id), function (status, response) {
        let previewInfo = document.getElementById("previewInfo");
        let previewItems = document.getElementById("previewItems");
        previewItems.innerHTML = "";
        previewInfo.innerHTML = response.Position + "/" + response.TotalCount;

        response.Items.map(function (value, index, array) {
            let a = document.createElement("a");
            a.classList.add("image");
            a.classList.add("position" + index);
            a.href = createUrl("/play/", value.Id, channelId);

            var img = new Image(value.Thumbnail.Width);
            img.src = value.Thumbnail.Url;

            a.appendChild(img);
            previewItems.appendChild(a);
        })

    })
}


function getJSON(url, callback) {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);
    xhr.responseType = 'json';
    xhr.onload = function () {
        var status = xhr.status;
        if (status === 200) {
            callback(null, xhr.response);
        } else {
            callback(status, xhr.response);
        }
    };
    xhr.send();
}

let nightModeCurrent = 0.0;
let opacityInc = 0.05;
let opacityTimer;
let nightModeDuration = 60 * 20;
let nightModeEnabled = false;

function startNightMode() {
    nightModeEnabled = true;
    opacityTimer = window.setInterval(function () {
        let overlay = document.getElementById("playerOverlay");
        nightModeCurrent += opacityInc;
        overlay.style.opacity = nightModeCurrent;
        player.setVolume(100 - (100 * nightModeCurrent));
        if (nightModeCurrent >= 1) {
            clearInterval(opacityTimer);
            player.pauseVideo();
        }
    }, (1000 * nightModeDuration) / (1 / opacityInc));
}

function stopNightMode() {
    nightModeEnabled = false;
    let overlay = document.getElementById("playerOverlay");
    overlay.style.opacity = 0.0;
    nightModeCurrent = 0.0;
    player.setVolume(100);
    clearInterval(opacityTimer);
}

function registerNightModeIcon() {
    document.getElementById("sleepIcon").addEventListener("click", function () {
        if (nightModeEnabled) {
            stopNightMode();
        } else {
            startNightMode();
        }
    });

}