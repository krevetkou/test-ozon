insert into comments (id, post_id, created_at, comment, comment_id)
values (4, '1', '2024-06-13T00:35:36.5502184+03:00', 'Mda', 4);

insert into comments (id, post_id, created_at, comment, comment_id)
values (5, '', '2024-06-13T00:35:36.5502184+03:00', 'Mda', 4);

insert into posts (id, title, text, created_at, can_be_commented, comment_id)
values (1, 'Help me', 'Help me please', '2024-04-13T00:35:36.5502184+03:00', true, 4);

insert into posts (id, title, text, created_at, can_be_commented, comment_id)
values (2, 'Hello', 'Hello world', '2024-05-13T00:35:36.5502184+03:00', false, 4);

insert into posts (id, title, text, created_at, can_be_commented, comment_id)
values (3, 'Here is', 'Hear is a cat', '2024-06-13T00:35:36.5502184+03:00', true, 4);

