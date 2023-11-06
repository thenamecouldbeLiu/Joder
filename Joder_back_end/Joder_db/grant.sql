GRANT ALL PRIVILEGES ON user_info.* TO 'online_user'@'localhost';
grant alter,create,delete,drop,index,insert,select,update,trigger,alter routine,
create routine, execute, create temporary tables on user_info.* to 'online_user';

create ROLE developer_user;
grant alter,create,delete,drop,index,insert,select,update,trigger,alter
 routine,create routine, execute, create temporary tables 
on user_info.* to 'developer_user';

create ROLE common_user;
grant delete,drop,insert,select,update, execute
on user_info.* to 'common_user';

grant 'common_user' to 'online_user'@'localhost';

select * from user;