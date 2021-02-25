CREATE TABLE public.moeobject_dbfile (
                                         file_id varchar(100) NOT NULL,
                                         metadata varchar(4000) NOT NULL,
                                         filecontent bytea NOT NULL,
                                         time_created bigint NOT NULL,
                                         time_updated bigint NOT NULL,
                                         CONSTRAINT moeobject_dbfile_pk PRIMARY KEY (file_id)
);
