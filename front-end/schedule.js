var sample = JSON.parse('{"refreshedAt":"25may2020 10:00pm","lastUpdate":"25may2020 9:30pm","times":[{"id":1,"start":"6:30pm","end":"6:55pm"},{"id":2,"start":"7:00pm","end":"7:25pm"},{"id":3,"start":"7:30pm","end":"7:55pm"},{"id":4,"start":"8:00pm","end":"8:25pm"},{"id":5,"start":"8:30pm","end":"8:55pm"},{"id":6,"start":"9:00pm","end":"9:25pm"},{"id":7,"start":"","end":""}],"rows":[{"room":"120","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]},{"room":"130","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]},{"room":"140","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]},{"room":"150","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]},{"room":"160","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]},{"room":"170","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]}]}');

$(function() {
    // var url = "/api/v1/schedule"
    // httpGetAsync(url, loadSchedule)
    loadSchedule(sample)
});

function loadSchedule(data) {
    var table = document.getElementById("schedule");
    var headerRow = document.getElementById("header-row");

    for (var slot of data["times"]) {
        var cell = document.createElement('th');
        cell.innerHTML = "<span>" + slot.start + " - " + slot.end + "</span>";
        headerRow.appendChild(cell);
    }
    table.appendChild(headerRow);

    for (var track of data["rows"]) {
        var row = document.createElement('tr');
        var labelCell = document.createElement('td');
        labelCell.innerHTML = "<span>" + track['room'] + "</span>";
        row.appendChild(labelCell);

        if (track["sessions"] != null) {
            for (var session of track["sessions"]) {
                var cell = document.createElement('td');
                cell.innerHTML = "<span class='title'>" + session.title + "</span><span class='speaker'>" + session.speaker + "</span>";
                row.appendChild(cell);
            }
        }
        table.appendChild(row);
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
