-- this is just sample data for testing, always drop and recreate database first so the IDs match

-- hashed password is '123456'
insert into users (name, nick, email, password)
values
('Felipe Reis', 'blobby', 'mail@mail.com', '$2a$10$ccQBnh020YczQ4iFwGRSQ.oRBKsrMXmW4fuw/62EiKLwpDLWTGoYy'),
('Felipe Reis 2', 'blobby2', 'mail2@mail.com', '$2a$10$ccQBnh020YczQ4iFwGRSQ.oRBKsrMXmW4fuw/62EiKLwpDLWTGoYy'),
('Felipe Reis 3', 'blobby3', 'mail3@mail.com', '$2a$10$ccQBnh020YczQ4iFwGRSQ.oRBKsrMXmW4fuw/62EiKLwpDLWTGoYy');

insert into followers (user_id, follower_user_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publications (title, content, author_id)
values
("Publication from user 1", "This is the publication from user 1: blobby", 1),
("Publication from user 2", "This is the publication from user 2: blobby2", 2),
("Publication from user 3", "This is the publication from user 3: blobby3", 3);