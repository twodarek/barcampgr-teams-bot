CREATE DATABASE IF NOT EXISTS barcampgr_schedule;

USE barcampgr_schedule;

CREATE TABLE IF NOT EXISTS Sessions (
    sessionID INTEGER NOT NULL AUTO_INCREMENT,
    title VARCHAR(256) NOT NULL,
    speaker VARCHAR(256) NOT NULL,
    deleted BOOL DEFAULT FALSE,
    timeUpdated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (sessionID)
);

CREATE TABLE IF NOT EXISTS Rooms (
    roomID INTEGER NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL,
    PRIMARY KEY (roomID)
);

CREATE TABLE IF NOT EXISTS Times (
     timeID INTEGER NOT NULL AUTO_INCREMENT,
     start VARCHAR(64) NOT NULL,
     end VARCHAR(64) NOT NULL,
     active BOOL NOT NULL DEFAULT False,
     PRIMARY KEY (timeID)
);

CREATE TABLE IF NOT EXISTS Schedule (
     relationID INTEGER NOT NULL AUTO_INCREMENT,
     timeID INTEGER NOT NULL,
     roomID INTEGER NOT NULL,
     sessionID INTEGER,
     timeUpdated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     active BOOL DEFAULT True,
     FOREIGN KEY (timeID) REFERENCES Times(timeID),
     FOREIGN KEY (roomID) REFERENCES Rooms(roomID),
     FOREIGN KEY (sessionID) REFERENCES Sessions(sessionID),
     PRIMARY KEY (relationID)
);

INSERT INTO Rooms (name) VALUES ('Main Room'), ('120'), ('130'), ('140'), ('150'), ('160'), ('170');
INSERT INTO Times (start, end) VALUES ('6:30pm', '6:55pm'), ('7:00pm', '7:25pm'), ('7:30pm', '7:55pm'), ('8:00pm', '8:25pm'), ('8:30pm', '8:55pm'), ('9:00pm', '9:25pm'), ('9:30pm', '9:55pm');