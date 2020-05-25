var sample = JSON.parse('{"refreshedAt": "25may2020 10:00pm", "lastUpdate": "25may2020 9:30pm", "times": [{      "id": 1,      "start": "6:30pm",      "end": "6:55pm"    },    {      "id": 2,      "start": "7:00pm",      "end": "7:25pm"    },    {      "id": 3,      "start": "7:30pm",      "end": "7:55pm"    },    {      "id": 4,      "start": "8:00pm",      "end": "8:25pm"    },    {      "id": 5,      "start": "8:30pm",      "end": "8:55pm"    },    {      "id": 6,      "start": "9:00pm",      "end": "9:25pm"    },    {      "id": 7,      "start": "",      "end": ""    }  ],  "rows": [    {      "room": "120",      "sessions": [        {          "time": 1,          "title": "This is a great talk",          "speaker": "Thomas Wodarek"        },        {          "time": 2,          "title": "This is a great talk",          "speaker": "Thomas Wodarek"        },        {          "time": 3,          "title": "This is a great talk",          "speaker": "Thomas Wodarek"        },        {          "time": 4,          "title": "This is a great talk",          "speaker": "Thomas Wodarek"        },        {          "time": 5,          "title": "This is a great talk",          "speaker": "Thomas Wodarek"        },        {          "time": 6,          "title": "This is a great talk",          "speaker": "Thomas Wodarek"        },        {          "time": 7,          "title": "This is a great talk",          "speaker": "Thomas Wodarek"        }      ]    }  ]}');

$(function() {
    loadSchedule(sample);
});

function loadSchedule(data) {
    var table = document.getElementById("schedulee");
    var headerRow = document.createElement('tr');
    var logoCell = document.createElement('th');
    logoCell.innerHTML = "<span>" + "BarCamp GR" + "</span>";
    headerRow.appendChild(logoCell);

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

        for (var session of track["sessions"]) {
            var cell = document.createElement('td');
            cell.innerHTML = "<span class='title'>" + session.title + "</span><span class='speaker'>" + session.speaker + "</span>";
            row.appendChild(cell);
        }
        table.appendChild(row);
    }
}
