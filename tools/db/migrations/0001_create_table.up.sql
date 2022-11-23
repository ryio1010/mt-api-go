CREATE SCHEMA IF NOT EXISTS muscle_tracking_go AUTHORIZATION mtadmin;

CREATE TABLE IF NOT EXISTS muscle_tracking_go.m_user(
    "userid" VARCHAR(255) NOT NULL,
    "username" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "regid" VARCHAR(255),
    "regdate" TIMESTAMP(0) WITHOUT TIME ZONE,
    "updid" VARCHAR(255),
    "upddate" TIMESTAMP(0) WITHOUT TIME ZONE,
    "version" INTEGER NOT NULL
);

ALTER TABLE
    muscle_tracking_go.m_user ADD PRIMARY KEY("userid");

CREATE TABLE IF NOT EXISTS  muscle_tracking_go.m_musclepart(
    "musclepartid" VARCHAR(255) NOT NULL,
    "musclepartname" VARCHAR(255) NOT NULL,
    "regid" VARCHAR(255),
    "regdate" TIMESTAMP(0) WITHOUT TIME ZONE,
    "updid" VARCHAR(255),
    "upddate" TIMESTAMP(0) WITHOUT TIME ZONE,
    "version" INTEGER NOT NULL
);

ALTER TABLE
    muscle_tracking_go.m_musclepart ADD PRIMARY KEY("musclepartid");

CREATE TABLE IF NOT EXISTS  muscle_tracking_go.m_menu(
    "menuid" INTEGER NOT NULL,
    "menuname" VARCHAR(255) NOT NULL,
    "musclepartid" VARCHAR(255) NOT NULL,
    "userid" VARCHAR(255) NULL,
    "status" VARCHAR(255) NOT NULL DEFAULT '1',
    "regid" VARCHAR(255),
    "regdate" TIMESTAMP(0) WITHOUT TIME ZONE,
    "updid" VARCHAR(255),
    "upddate" TIMESTAMP(0) WITHOUT TIME ZONE,
    "version" INTEGER NOT NULL
);

ALTER TABLE
    muscle_tracking_go.m_menu ADD PRIMARY KEY("menuid");

COMMENT
ON COLUMN
    muscle_tracking_go.m_menu."status" IS '0:デフォルト
1:ユーザー追加';

CREATE TABLE IF NOT EXISTS  muscle_tracking_go.t_traininglog(
    "logid" INTEGER NOT NULL,
    "menuid" INTEGER NOT NULL,
    "trainingweight" DOUBLE PRECISION NOT NULL,
    "trainingcount" INTEGER NOT NULL,
    "trainingdate" VARCHAR(255) NOT NULL,
    "trainingmemo" VARCHAR(255) NOT NULL,
    "userid" VARCHAR(255) NOT NULL,
    "regid" VARCHAR(255),
    "regdate" TIMESTAMP(0) WITHOUT TIME ZONE,
    "updid" VARCHAR(255),
    "upddate" TIMESTAMP(0) WITHOUT TIME ZONE,
    "version" INTEGER NOT NULL
);

ALTER TABLE
    muscle_tracking_go.t_traininglog ADD PRIMARY KEY("logid");

CREATE TABLE IF NOT EXISTS  muscle_tracking_go.t_bodycomp(
    "userid" VARCHAR(255) NOT NULL,
    "trainingdate" VARCHAR(255) NOT NULL,
    "height" DOUBLE PRECISION NOT NULL,
    "weight" DOUBLE PRECISION NOT NULL,
    "bfp" DOUBLE PRECISION NOT NULL,
    "regid" VARCHAR(255),
    "regdate" TIMESTAMP(0) WITHOUT TIME ZONE,
    "updid" VARCHAR(255),
    "upddate" TIMESTAMP(0) WITHOUT TIME ZONE,
    "version" INTEGER NOT NULL
);

ALTER TABLE
    muscle_tracking_go.t_bodycomp ADD PRIMARY KEY("userid","trainingdate");

ALTER TABLE
    muscle_tracking_go.m_menu ADD CONSTRAINT "m_menu_userid_foreign" FOREIGN KEY("userid") REFERENCES "m_user"("userid");

ALTER TABLE
    muscle_tracking_go.t_traininglog ADD CONSTRAINT "t_traininglog_userid_foreign" FOREIGN KEY("userid") REFERENCES "m_user"("userid");

ALTER TABLE
    muscle_tracking_go.t_traininglog ADD CONSTRAINT "t_traininglog_menuid_foreign" FOREIGN KEY("menuid") REFERENCES "m_menu"("menuid");

ALTER TABLE
    muscle_tracking_go.m_menu ADD CONSTRAINT "m_menu_musclepartid_foreign" FOREIGN KEY("musclepartid") REFERENCES "m_musclepart"("musclepartid");

CREATE SEQUENCE IF NOT EXISTS  m_menu_menuid_seq;
ALTER TABLE muscle_tracking_go.m_menu ALTER COLUMN "menuid" SET DEFAULT nextval('m_menu_menuid_seq');

CREATE SEQUENCE IF NOT EXISTS  t_traininglog_logid_seq;
ALTER TABLE muscle_tracking_go.t_traininglog ALTER COLUMN "logid" SET DEFAULT nextval('t_traininglog_logid_seq');
