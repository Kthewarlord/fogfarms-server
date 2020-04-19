—Populate data
Insert into location (city,province) VALUES ('bkk','bkk');
Insert into location (locationid,city,province) VALUES (0,'null','null');
Insert into nutrient (nutrientid,Part,Nitrogen,Phosphorus,Potassium) Values(0,0,0,0,0);
Insert into nutrient (Part,Nitrogen,Phosphorus,Potassium) Values (1,1,1,1);
Insert into plant (plantid,name,tds,ph,lux,lightsonhour,lightsoffhour) Values (0,'null' ,0,0,0,0,0);
Insert into plant (name,tds,ph,lux,lightsonhour,lightsoffhour) values ('plant' ,0,0,0,0,0);

Insert into ModuleGroup (modulegroupid, modulegrouplabel,plantid,locationid,param_tds,param_ph,param_humidity,OnAuto,Lightsonhour,lightsoffhour,timerlastreset)
values (0,'null',0,0,0,0,0,false,0,0,now());
Insert into ModuleGroup (modulegrouplabel,plantid,locationid,param_tds,param_ph,param_humidity,OnAuto,Lightsonhour,lightsoffhour,timerlastreset)
values ('modulegroup1',1,1,3,3,3,true,2,22,now());
Insert into ModuleGroup (modulegrouplabel,plantid,locationid,param_tds,param_ph,param_humidity,OnAuto,Lightsonhour,lightsoffhour,timerlastreset)
values ('modulegroup2',1,1,3,3,3,true,2,22,now());

Insert into module(moduleid,modulegroupid,modulelabel,token) Values(0,0,'null','null');
Insert into module(modulegroupid,modulelabel,token) Values(1,'module1','module1');
Insert into module(modulegroupid,modulelabel,token) Values(1,'module2','module2');
Insert into module(modulegroupid,modulelabel,token) Values(2,'module3','module3');

Insert into nutrientunit(nutrientunitid,moduleid,nutrientid) Values(0,0,0);
Insert into nutrientunit(moduleid,nutrientid) Values(1,1);
Insert into nutrientunit(moduleid,nutrientid) Values(2,1);

Insert into growunit(growunitid,moduleid,capacity) Values(0,0,0);
Insert into growunit(moduleid,capacity) Values(1,1);
Insert into growunit(moduleid,capacity) Values(2,1);
Insert into growunit(moduleid,capacity) Values(1,1);

insert into phupunit (phupunitid, nutrientunitid) Values(0,0);
insert into phupunit (nutrientunitid) Values(1);
insert into phupunit (nutrientunitid) Values(2);

insert into phdownunit (phdownunitid, nutrientunitid) Values(0,0);
insert into phdownunit (nutrientunitid) Values(1);
insert into phdownunit (nutrientunitid) Values(2);

insert into devicetype(Devicetypeid,devicetype) Values(0,'null');
insert into devicetype(devicetype) Values('solenoidValve');

insert into device(deviceid,devicetypeid,ison,growunitid,nutrientunitid,phdownunitid,phupunitid) Values(0,0,false,0,0,0,0);
insert into device(devicetypeid,ison,growunitid,nutrientunitid,phdownunitid,phupunitid) Values(1,false,1,0,0,0);
insert into device(devicetypeid,ison,growunitid,nutrientunitid,phdownunitid,phupunitid) Values(1,false,0,1,0,0);

insert into Sensordata(moduleid,timestamp,tds,ph,solutiontemperature,ARRGrowunitlux,arrgrowunithumidity,arrgrowunittemperature) Values(1,Now(),1,1,1,'{1,2}','{1,3}','{2,3}');
insert into Sensordata(moduleid,timestamp,tds,ph,solutiontemperature,ARRGrowunitlux,arrgrowunithumidity,arrgrowunittemperature) Values(2,Now(),1,1,1,'{2}','{3}','{3}');