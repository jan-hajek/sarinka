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

function getChannelsUrl() {
    return url + "/channels/";
}