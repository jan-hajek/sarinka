let port = location.port;

let url = location.protocol + "//" + location.hostname;
if (port !== "80") {
    url += ":" + port
}
let urlParams = new URLSearchParams(location.search);
let position = 0;
let channels = [];
let channelId = urlParams.get("channelId");
let currentId = urlParams.get("id");
let nextId = null;
let player;
let screenWidth = screen.width;

function selectCurrentChannel() {
    let left = screenWidth / 2;

    left -= 120;

    left -= position * 190;

    document.getElementById("channels").style.left = left + "px";

    document.getElementById("channel" + position).classList.add("selected");
}

function registerHomepageOnKeydown() {
    document.onkeydown = function (event) {
        document.getElementById("channel" + position).classList.remove("selected");

        switch (event.key) {
            case "ArrowRight":
                position = Math.min(position + 1, channels.length - 1);
                break;
            case "ArrowLeft":
                position = Math.max(position - 1, 0);
                break;
            case "Enter":
                location.href = "play/?channelId=" + channels[position].Id;

                break;
        }
        selectCurrentChannel()
    };
}

function getChannelsUrl() {
    return url + "/channels/";
}


function startHomepage() {
    registerHomepageOnKeydown();

    getJSON(getChannelsUrl(), function (status, response) {
        let channelsDiv = document.getElementById("channels");
        channelsDiv.innerHTML = "";

        channels = response.Channels;
        channelsDiv.style.width = (250 * response.Channels.length) + "px";

        response.Channels.map(function (value, index, array) {
            var div = document.createElement("div");
            div.classList.add("channel");
            if (value.Id == channelId) {
                position = index
            }
            div.id = "channel" + index;

            var img = new Image();
            img.src = value.ThumbnailUrl;
            img.classList.add("img");

            var text = document.createElement("div");
            text.innerHTML = value.Name;
            text.classList.add("title");

            div.appendChild(img);
            channelsDiv.appendChild(div);
        });


        selectCurrentChannel();

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
    currentId = response.Current.Id;
    nextId = response.Next.Id;
    document.title = response.Current.Title;
    setYoutubeId(currentId);
    loadPreview(nextId);
}

function setYoutubeId(id) {
    if (!player) {
        let width = screenWidth - 280;
        let height = screen.height - 200;

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
    if (event.data == YT.PlayerState.ENDED) {
        nextVideo();
    }
    if (event.data == YT.PlayerState.PLAYING) {
        playing = true
    } else {
        playing = false;
    }
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
                if (channelId !== null) {
                    location.href = url + "?channelId=" + channelId;
                } else {
                    location.href = url;
                }
                break;
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
            var div = document.createElement("div");
            div.classList.add("image");
            div.classList.add("position" + index);

            var img = new Image(value.Thumbnail.Width);
            img.src = value.Thumbnail.Url;

            div.appendChild(img);
            previewItems.appendChild(div);
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