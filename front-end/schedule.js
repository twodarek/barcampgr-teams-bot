var sample = JSON.parse('{"refreshedAt":"25may2020 10:00pm","lastUpdate":"25may2020 9:30pm","times":[{"id":1,"start":"6:30pm","end":"6:55pm"},{"id":2,"start":"7:00pm","end":"7:25pm"},{"id":3,"start":"7:30pm","end":"7:55pm"},{"id":4,"start":"8:00pm","end":"8:25pm"},{"id":5,"start":"8:30pm","end":"8:55pm"},{"id":6,"start":"9:00pm","end":"9:25pm"},{"id":7,"start":"","end":""}],"rows":[{"room":"120","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]},{"room":"130","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]},{"room":"140","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]},{"room":"150","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]},{"room":"160","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]},{"room":"170","sessions":[{"time":1,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":2,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":3,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":4,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":5,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":6,"title":"This is a great talk","speaker":"Thomas Wodarek"},{"time":7,"title":"This is a great talk","speaker":"Thomas Wodarek"}]}]}');
var sample2 = JSON.parse('{"refreshedAt":"","lastUpdate":"","times":[{"id":1,"start":"5:00pm","end":"5:55pm","day":"Friday","displayable":true},{"id":2,"start":"6:00pm","end":"6:25pm","day":"Friday","displayable":true},{"id":3,"start":"6:30pm","end":"6:55pm","day":"Friday","displayable":true},{"id":4,"start":"7:00pm","end":"7:25pm","day":"Friday","displayable":true},{"id":5,"start":"7:30pm","end":"7:55pm","day":"Friday","displayable":true},{"id":6,"start":"8:00pm","end":"8:25pm","day":"Friday","displayable":true},{"id":7,"start":"8:30pm","end":"8:55pm","day":"Friday","displayable":true}],"rows":[{"room":"Main Room","sessions":[{"time":6,"room":1,"title":"derpity derp derp","speaker":"Derpy McDerperson"}]},{"room":"120","sessions":[{"time":3,"room":2,"title":"DERP","speaker":"Thomas Wodarek"}]},{"room":"130","sessions":null},{"room":"140","sessions":[{"time":7,"room":4,"title":"This is an awesome talk","speaker":"thomas wodarek"}]},{"room":"150","sessions":null},{"room":"160","sessions":[{"time":2,"room":6,"title":"idk","speaker":"meh"},{"time":1,"room":6,"title":"This is a mediocre talk","speaker":"Some Shmuck"}]},{"room":"170","sessions":null}]}');
var sample3 = JSON.parse('{"refreshedAt":"","lastUpdate":"","times":[{"id":1,"start":"6:30pm","end":"6:55pm","day":"Friday","displayable":true},{"id":2,"start":"7:00pm","end":"7:25pm","day":"Friday","displayable":true},{"id":3,"start":"7:30pm","end":"7:55pm","day":"Friday","displayable":true},{"id":4,"start":"8:00pm","end":"8:25pm","day":"Friday","displayable":true},{"id":5,"start":"8:30pm","end":"8:55pm","day":"Friday","displayable":true},{"id":6,"start":"9:00pm","end":"9:25pm","day":"Friday","displayable":true}],"rows":[{"room":"General","sessions":[{"time":"1","room":"1","title":"Opening Session","speaker":"Organizers","uniqueString":"2efxFLP2SaH2m3Lh8QsF2AYkkBLsvBkQXDCNJVbLVJWilCN5VPh9mFG0alU4OIBo","altText":"Opening Session by Organizers in General at 6:30pm"}]},{"room":"2020","sessions":[{"time":"1","room":"2","title":"Blocked","speaker":"","uniqueString":"ZLH9V00yOyTwg0mSMu1Eq8Wg0FUGFv4aIDttghpgiXAvZXEwq7a0LOTI0pHGIyoG","altText":"Unscheduleable time"}]},{"room":"Creative Corner","sessions":[{"time":"2","room":"3","title":"foobar","speaker":"Thomas Wodarek","uniqueString":"k3hkwde2wPmWiqUwdIEkBxnVrwq04NsBHcUAqs0Xrk2n188c2NH9S3NJCCZmjvqO","altText":"foobar by Thomas Wodarek in Creative Corner at 7:00pm"},{"time":"4","room":"3","title":"testtest2","speaker":"Thomas Wodarek","uniqueString":"4AP3MZPMPLdakrGsaQ5KzkEpKjQDhA8T9m7EpkY2VFAKIiQNOTn4k0NQrjKipwvJ","altText":"testtest2 by Thomas Wodarek in Creative Corner at 8:00pm"},{"time":"1","room":"3","title":"Blocked","speaker":"","uniqueString":"jxXu2JvOntEgo3TxvuxESkCP5z5GL4sYMmM11hOnAe8rgKLdjG0ZETuwgvd58KFR","altText":"Unscheduleable time"}]},{"room":"Programming","sessions":[{"time":"5","room":"4","title":"talkity talk talk","speaker":"Thomas Wodarek","uniqueString":"87xx6XTm0iRHA6m0OSPhYLOpzrMXb7hph19zHYjdTS5iDdKNbxWuJb8mGIj55Qdl","altText":"talkity talk talk by Thomas Wodarek in Programming at 8:30pm"},{"time":"4","room":"4","title":"foo bar baz","speaker":"Dave Brondsema","uniqueString":"eAWW713lUu2Sif0AEmLQ24ra1gJP8E9gSn9hdzTvcFfCTQHKt6eAx0S4wzcLpaFL","altText":"foo bar baz by Dave Brondsema in Programming at 8:00pm"},{"time":"3","room":"4","title":"This is a new test","speaker":"Thomas Wodarek","uniqueString":"5pYHT8NQI0hc62XoRYKnUwBqlSMTp7ns1XV6FIxkvdytXC20XcuU0xtMGwAPa2pL","altText":"This is a new test by Thomas Wodarek in Programming at 7:30pm"},{"time":"1","room":"4","title":"Blocked","speaker":"","uniqueString":"RrIy36yKPDk4sPZYdgmvY2g6K94xoDca0uh51r83zBmC35Zm0T21kKBPa3bjS6GY","altText":"Unscheduleable time"}]},{"room":"Room 120","sessions":[{"time":"1","room":"5","title":"Blocked","speaker":"","uniqueString":"bmTmpuRnyDb8CVpCDilEUufOxJ5SG99YstgG5m0vOg0LvgNztFJsEKGsF3a1HnaL","altText":"Unscheduleable time"}]},{"room":"Room 140","sessions":[{"time":"3","room":"6","title":"test talk","speaker":"Thomas Wodarek","uniqueString":"LcET7UuT0Jb2vcsV7TD0WzHUPhnhI2MdUn4OE4gRuQ78FfhPV1LUPweQgBHM3vDB","altText":"test talk by Thomas Wodarek in Room 140 at 7:30pm"},{"time":"4","room":"6","title":"test talk","speaker":"Thomas Wodarek","uniqueString":"rtXwMr4V63MuIqDFv4c9DLRKrZ5VhSK1D2RcSjqHsmAustAfEEar5MvTNsbmCMFV","altText":"test talk by Thomas Wodarek in Room 140 at 8:00pm"},{"time":"1","room":"6","title":"Blocked","speaker":"","uniqueString":"4g1sLn4xfsm3FWjcXLt3qybJiQX2xwLS2JhEGrfFwi5Fv2491rNdeSiJ9zcRA1QW","altText":"Unscheduleable time"},{"time":"2","room":"6","title":"test talk 2","speaker":"Thomas Wodarek","uniqueString":"t39L8buoFWA5S8fLA3RInJFNrwhmpjJnNVwai5p8ezTMehRlBtJGlZR4AbWhFj2e","altText":"test talk 2 by Thomas Wodarek in Room 140 at 7:00pm"}]},{"room":"Room 170","sessions":[{"time":"6","room":"7","title":"Pull a rabbit from a hat for 500","speaker":"Nicole Scheffler","uniqueString":"SKfT3aqvMMz1KHX2lFJA18WtdCDaUlkMWkSGNq4EjtTbRmQlG3AYR5P5dKOUkMZX","altText":"Pull a rabbit from a hat for 500 by Nicole Scheffler in Room 170 at 9:00pm"},{"time":"1","room":"7","title":"Blocked","speaker":"","uniqueString":"6yiuhRhEDgSRkPYzs34fLrFgV6Buts5HDdzoNtMJrDKZF0R0f46taC7cU5dMUeLJ","altText":"Unscheduleable time"}]},{"room":"Wellness","sessions":[{"time":"3","room":"8","title":"This is another test","speaker":"Thomas Wodarek","uniqueString":"c2T1zXOpso8K1QxDekqzJfM1E4B3aEi5AGWu3X5oCKoNa831wm4HOVPxVMRgkhyb","altText":"This is another test by Thomas Wodarek in Wellness at 7:30pm"},{"time":"1","room":"8","title":"Blocked","speaker":"","uniqueString":"QKFVh85SN6UDhPb9F1tRFtvUe8ULHSvkI0V7qyHsRj4jJA31yy4zyxCexmP5cFwK","altText":"Unscheduleable time"},{"time":"4","room":"8","title":"Heres a great talk idea","speaker":"Thomas Wodarek","uniqueString":"B7Tdo6uyIMpIiSKWsaCiRSdytIeWDCGHWRcR6Nhdmli96Qhk4fJKoZdbIhU5d3yu","altText":"Heres a great talk idea by Thomas Wodarek in Wellness at 8:00pm"},{"time":"6","room":"8","title":"fawuehflawhe","speaker":"Thomas Wodarek","uniqueString":"7iRjGLZjQu1a2gnlscZWActRSVPdt61nGEVnzkER2klgfVjVPYmmmI9PftnBi4td","altText":"fawuehflawhe by Thomas Wodarek in Wellness at 9:00pm"}]}]}');

$(function() {
    var url = "/api/v1/schedule";
    // httpGetAsync(url, createSchedule);
    createSchedule(sample3);
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
            cell.innerHTML = "<img src='transparent-pixel.png' class='a11y-pxl' alt='empty slot'>";
            row.appendChild(cell);
        }
        table.appendChild(row);
    }
    populateSchedule(data);
}

function populateSchedule(data) {
    for (var track of data["rows"]) {
        if (track["sessions"] != null) {
            for (var session of track["sessions"]) {
                var targetCellSelector = '#' + getSlotId(track['room'], session['time']);
                $(targetCellSelector).html($(targetCellSelector).html() + "<span class='title'>" + session.title + "</span><span class='speaker'>" + session.speaker + "</span>");
                $(targetCellSelector).find(".a11y-pxl").attr("alt", session.altText);
                $(targetCellSelector).attr("aria-description", session.altText);
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
