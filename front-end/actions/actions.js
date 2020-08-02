function loadPage() {
    var sessionStr = getParameterByName('unique_str');
    var callUrl = "/api/v1/session/"+sessionStr;
    httpGetAsync(callUrl, loadSession);
}

function loadSession(res) {
    if (res === "{}") {
        document.getElementById('main').align = 'center';
        document.getElementById('main').innerText = "Session Not Found";
        return
    }
    var sessionData = JSON.parse(res);
    document.getElementById('title').value = sessionData.title;
    document.getElementById('speaker').value = sessionData.speaker;
    window.sessionData = sessionData


    var timesUrl = "/api/v1/times";
    httpGetAsync(timesUrl, loadTimes);

    var roomsUrl = "/api/v1/rooms";
    httpGetAsync(roomsUrl, loadRooms);

    document.getElementById("save-button").addEventListener("click", updateSession)
}

function loadTimes(res) {
    var timesData = JSON.parse(res);
    if (timesData == {}) {
        return
    }
    var selector = document.getElementById('timeSelector')

    timesData.forEach(function(time) {
        var option = document.createElement("option");
        option.label = time.start;
        option.value = time.id;
        option.innerText = time.start;
        if (time.id == window.sessionData.time) {
            option.selected = true;
        }
        selector.add(option);
    })
}

function loadRooms(res) {
    var roomsData = JSON.parse(res);
    if (roomsData == {}) {
        return
    }
    var selector = document.getElementById('roomSelector')

    roomsData.forEach(function(room) {
        var option = document.createElement("option");
        option.label = room.name;
        option.value = room.id;
        option.innerText = room.name;
        if (room.id == window.sessionData.room) {
            option.selected = true;
        }
        selector.add(option);
    })
}

function eatSubmit(e) {
    console.log("running")
    e.preventDefault();
    return false
}

function updateSession(e) {
    var sessionData = window.sessionData;
    var updated = false;
    if (sessionData.title !== document.getElementById('title').value) {
        sessionData.title = document.getElementById('title').value;
        updated = true;
    }
    if (sessionData.speaker !== document.getElementById('speaker').value) {
        sessionData.speaker = document.getElementById('speaker').value;
        updated = true;
    }
    var roomSelector = document.getElementById('roomSelector')
    if (sessionData.room !== roomSelector.options[roomSelector.selectedIndex].value) {
        sessionData.room = roomSelector.options[roomSelector.selectedIndex].value;
        updated = true;
    }
    var timeSelector = document.getElementById('timeSelector')
    if (sessionData.time !== timeSelector.options[timeSelector.selectedIndex].value) {
        sessionData.time = timeSelector.options[timeSelector.selectedIndex].value;
        updated = true
    }

    if (updated) {
        var sessionStr = getParameterByName('unique_str');
        var postUrl = "/api/v1/session/" + sessionStr
        httpPostAsync(postUrl, updateSuccessful, updateFailed, JSON.stringify(sessionData));
    } else {
        alert("No changes to save")
    }
}

function updateSuccessful() {
    alert("Update successful");
    window.location.href = "https://talks.barcampgr.org/"
}

function updateFailed(response) {
    alert("Update Failed because "+response);
}

// https://stackoverflow.com/questions/247483/http-get-request-in-javascript
function httpGetAsync(theUrl, callback) {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function() {
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200)
            callback(xmlHttp.responseText);
    }
    xmlHttp.open("GET", theUrl, true); // true for asynchronous
    xmlHttp.send(null);
}

function httpPostAsync(theUrl, successCallback, failCallback, body) {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function() {
        if(xmlHttp.readyState === XMLHttpRequest.DONE) {
            var status = xmlHttp.status;
            if (status === 0 || (status >= 200 && status < 400)) {
                successCallback();
            } else {
                failCallback(xmlHttp.responseText);
            }
        }
    }
    xmlHttp.open("POST", theUrl, true);
    xmlHttp.send(body);
}


// https://stackoverflow.com/questions/901115/how-can-i-get-query-string-values-in-javascript?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa
function getParameterByName(name, url) {
    if (!url) url = window.location.href;
    name = name.replace(/[\[\]]/g, "\\$&");
    var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
        results = regex.exec(url);
    if (!results) return null;
    if (!results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, " "));
}
