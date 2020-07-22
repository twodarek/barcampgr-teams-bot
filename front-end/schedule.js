var sample = JSON.parse('{"refreshedAt":"25may2020 10:00pm","lastUpdate":"25may2020 9:30pm","times":[{"id":1,"start":"6:30pm","end":"6:55pm"},{"id":2,"start":"7:00pm","end":"7:25pm"},{"id":3,"start":"7:30pm","end":"7:55pm"},{"id":4,"start":"8:00pm","end":"8:25pm"},{"id":5,"start":"8:30pm","end":"8:55pm"},{"id":6,"start":"9:00pm","end":"9:25pm"},{"id":7,"start":"","end":""}],"rows":[{"room":"120","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]},{"room":"130","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]},{"room":"140","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]},{"room":"150","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]},{"room":"160","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]},{"room":"170","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]}]}');
var sample2 = JSON.parse('{"refreshedAt":"","lastUpdate":"","times":[{"id":1,"start":"5:00pm","end":"5:55pm","day":"Friday","displayable":true},{"id":2,"start":"6:00pm","end":"6:25pm","day":"Friday","displayable":true},{"id":3,"start":"6:30pm","end":"6:55pm","day":"Friday","displayable":true},{"id":4,"start":"7:00pm","end":"7:25pm","day":"Friday","displayable":true},{"id":5,"start":"7:30pm","end":"7:55pm","day":"Friday","displayable":true},{"id":6,"start":"8:00pm","end":"8:25pm","day":"Friday","displayable":true},{"id":7,"start":"8:30pm","end":"8:55pm","day":"Friday","displayable":true}],"rows":[{"room":"Main Room","sessions":[{"time":6,"room":1,"title":"derpity derp derp","speaker":"Derpy McDerperson"}]},{"room":"120","sessions":[{"time":3,"room":2,"title":"DERP","speaker":"Thomas Wodarek"}]},{"room":"130","sessions":null},{"room":"140","sessions":[{"time":7,"room":4,"title":"This is an awesome talk","speaker":"thomas wodarek"}]},{"room":"150","sessions":null},{"room":"160","sessions":[{"time":2,"room":6,"title":"idk","speaker":"meh"},{"time":1,"room":6,"title":"This is a mediocre talk","speaker":"Some Shmuck"}]},{"room":"170","sessions":null}]}')

$(function() {
    var url = "/api/v1/schedule";
    httpGetAsync(url, createSchedule);
    // createSchedule(sample2);
});

function getSlotId(room, timeId) {
    var cellId = 'slot-' + room.toString().toLowerCase() + '-' + timeId.toString();
    const regex = / /gi;
    cellId = cellId.replace(regex, '-');
    return cellId;
}

function createSchedule(data) {
    var table = document.getElementById("schedule");
    var headerRow = document.getElementById("header-row");
    var timeIds = [];

    for (var slot of data["times"]) {
        var cell = document.createElement('th');
        timeIds.push(slot.id);
        cell.innerHTML = "<span>" + slot.start + " - " + slot.end + "</span>";
        headerRow.appendChild(cell);
    }
    table.appendChild(headerRow);

    for (var track of data["rows"]) {
        var row = document.createElement('tr');
        var labelCell = document.createElement('td');
        labelCell.innerHTML = "<span>" + track['room'] + "</span>";
        row.appendChild(labelCell);

        for (var id of timeIds) {
            var cell = document.createElement('td');
            cell.setAttribute('id', getSlotId(track['room'], id));
            row.appendChild(cell);
        }
        table.appendChild(row);
    }
    populateSchedule(data)
}

function populateSchedule(data) {
    for (var track of data["rows"]) {
        if (track["sessions"] != null) {
            for (var session of track["sessions"]) {
                var targetCellSelector = '#' + getSlotId(track['room'], session['time']);
                $(targetCellSelector).html("<span class='title'>" + session.title + "</span><span class='speaker'>" + session.speaker + "</span>");
            }
        }
    }
}

// https://stackoverflow.com/questions/247483/http-get-request-in-javascript
function httpGetAsync(theUrl, callback) {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function() {
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200)
            callback(JSON.parse(xmlHttp.responseText));
    }
    xmlHttp.open("GET", theUrl, true); // true for asynchronous
    xmlHttp.send(null);
}
