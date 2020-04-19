-- Execute this first

-- Plant
CREATE TABLE Plant (
    PlantID SERIAL NOT NULL,
    Name VARCHAR(256) UNIQUE NOT NULL,
    TDS FLOAT NOT NULL,
    PH FLOAT NOT NULL,
    Lux FLOAT NOT NULL,
    LightsOnHour FLOAT NOT NULL,
    LightsOffHour FLOAT NOT NULL,
    PRIMARY KEY (PlantID)
);

-- Location
CREATE TABLE Location (
    LocationID SERIAL NOT NULL,
    City VARCHAR(256) NOT NULL,
    Province VARCHAR(256) NOT NULL,
    PRIMARY KEY (LocationID),
    UNIQUE (City, Province)
);

-- Users
CREATE TABLE Users (
    UserID SERIAL NOT NULL,
    Username VARCHAR(256) UNIQUE NOT NULL,
    IsAdministrator BOOLEAN NOT NULL DEFAULT FALSE,
    Hash VARCHAR(256) NOT NULL,
    Salt VARCHAR(256) NOT NULL,
    CreatedAt TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (UserID)
);

-- Nutrient
CREATE TABLE Nutrient (
    NutrientID SERIAL NOT NULL,
    Part INT NOT NULL,
    Nitrogen INT NOT NULL,
    Phosphorus INT NOT NULL,
    Potassium INT NOT NULL,
    PRIMARY KEY (NutrientID),
    UNIQUE (Part, Nitrogen, Phosphorus, Potassium)
);

-- ModuleGroup
--get plantid for default parameters
--get locationid for information
CREATE TABLE ModuleGroup (
    ModuleGroupID SERIAL NOT NULL,
    ModuleGroupLabel VARCHAR(256) UNIQUE NOT NULL,
    PlantID INT NOT NULL,
    LocationID INT NOT NULL,
    Param_TDS FLOAT NOT NULL,
    Param_PH FLOAT NOT NULL,
    Param_Humidity FLOAT NOT NULL,
    OnAuto BOOLEAN NOT NULL,
    LightsOnHour FLOAT NOT NULL,
    LightsOffHour FLOAT NOT NULL,
    TimerLastReset TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (ModuleGroupID),
    FOREIGN KEY (PlantID) REFERENCES Plant (PlantID) ON DELETE SET DEFAULT ON UPDATE CASCADE,
    FOREIGN KEY (LocationID) REFERENCES Location (LocationID) ON DELETE SET DEFAULT ON UPDATE CASCADE
);

-- Module
CREATE TABLE Module (
    ModuleID SERIAL NOT NULL,
    ModuleGroupID INT NOT NULL DEFAULT 0,
    ModuleLabel VARCHAR(256) UNIQUE NOT NULL,
    Token VARCHAR(256) UNIQUE NOT NULL,
    PRIMARY KEY (ModuleID),
    FOREIGN KEY (ModuleGroupID) REFERENCES ModuleGroup (ModuleGroupID) ON DELETE SET DEFAULT ON UPDATE CASCADE
);
-- NutrientUnit
--get module id for the module it belongs to
--get nutrient id for nutrient composition
CREATE TABLE NutrientUnit (
    NutrientUnitID SERIAL NOT NULL,
    ModuleID INT NOT NULL,
    NutrientID INT NOT NULL,
    PRIMARY KEY (NutrientUnitID),
    FOREIGN KEY (ModuleID) REFERENCES Module (ModuleID) ON DELETE SET DEFAULT ON UPDATE CASCADE,
    FOREIGN KEY (NutrientID) REFERENCES Nutrient (NutrientID) ON DELETE SET DEFAULT ON UPDATE CASCADE
);

-- PHUpUnit
--get nutrientunitid it belongs to
CREATE TABLE PHUpUnit (
    PHUpUnitID SERIAL NOT NULL,
    NutrientUnitID INT NOT NULL,
    PRIMARY KEY (PHUpUnitID),
    FOREIGN KEY (NutrientUnitID) REFERENCES NutrientUnit (NutrientUnitID) ON DELETE SET DEFAULT ON UPDATE CASCADE
);

CREATE TABLE PHDownUnit (
    PHDownUnitID SERIAL NOT NULL,
    NutrientUnitID INT NOT NULL,
    PRIMARY KEY (PHDownUnitID),
    FOREIGN KEY (NutrientUnitID) REFERENCES NutrientUnit (NutrientUnitID) ON DELETE SET DEFAULT ON UPDATE CASCADE
);

-- SensorData
-- logs data from sensors for each module
CREATE TABLE SensorData (
    ModuleID INT NOT NULL,
    Timestamp TIMESTAMP NOT NULL DEFAULT NOW(),
    TDS FLOAT NOT NULL,
    PH FLOAT NOT NULL,
    SolutionTemperature FLOAT NOT NULL,
    ArrGrowUnitLux FLOAT ARRAY NOT NULL,
    ArrGrowUnitHumidity FLOAT ARRAY NOT NULL,
    ArrGrowUnitTemperature FLOAT ARRAY NOT NULL,
    PRIMARY KEY (Timestamp, ModuleID),
    FOREIGN KEY (ModuleID) REFERENCES Module (ModuleID) ON DELETE SET DEFAULT ON UPDATE CASCADE
);

-- GrowUnit
CREATE TABLE GrowUnit (
    GrowUnitID SERIAL NOT NULL,
    ModuleID INT NOT NULL,
    Capacity INT,
    PRIMARY KEY (GrowUnitID),
    FOREIGN KEY (ModuleID) REFERENCES Module (ModuleID) ON DELETE SET DEFAULT ON UPDATE CASCADE
);

CREATE TABLE DeviceType (
  DeviceTypeID SERIAL NOT NULL,
  DeviceType VARCHAR(256) NOT NULL,
  PRIMARY KEY (DeviceTypeID),
  UNIQUE (DeviceType)
);

-- Device
-- should delete device if it no longer exists
CREATE TABLE Device (
    DeviceID SERIAL NOT NULL,
    DeviceTypeID INT NOT NULL,
    IsOn BOOLEAN NOT NULL DEFAULT FALSE,
    GrowUnitID INT,
    NutrientUnitID INT,
    PHDownUnitID INT,
    PHUpUnitID INT,
    PRIMARY KEY (DeviceID),
    FOREIGN KEY (DeviceTypeID) REFERENCES DeviceType (DeviceTypeID) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (GrowUnitID) REFERENCES GrowUnit (GrowUnitID) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (NutrientUnitID) REFERENCES NutrientUnit (NutrientUnitID) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (PHDownUnitID) REFERENCES PHDownUnit (PHDownUnitID) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (PHUpUnitID) REFERENCES PHUpUnit (PHUpUnitID) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (DeviceTypeID, GrowUnitID, NutrientUnitID, PHDownUnitID, PHUpUnitID)
);

-- SensorData_ModuleGroup
--logs sensor data of module group
CREATE TABLE SensorData_ModuleGroup (
    ModuleGroupID INT NOT NULL,
    Timestamp timestamp NOT NULL DEFAULT NOW(),
    Humidity FLOAT NOT NULL,
    Temperature FLOAT NOT NULL,
    PRIMARY KEY (ModuleGroupID, Timestamp),
    FOREIGN KEY (ModuleGroupID) REFERENCES ModuleGroup (ModuleGroupID) ON DELETE SET DEFAULT ON UPDATE CASCADE
);

-- Permission
-- if user does not exist, permission does not need to be recorded
-- if module group id does not exist, permission can still be stored
CREATE TABLE Permission (
    UserID INT NOT NULL,
    ModuleGroupID INT NOT NULL,
    PermissionLevel INT NOT NULL DEFAULT 0,
    PRIMARY KEY (UserID, ModuleGroupID),
    FOREIGN KEY (UserID) REFERENCES Users (UserID) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (ModuleGroupID) REFERENCES ModuleGroup (ModuleGroupID) ON DELETE ON DEFAULT ON UPDATE CASCADE,
    UNIQUE (UserID, ModuleGroupID)
);