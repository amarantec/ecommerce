CREATE TABLE IF NOT EXISTS users (
  id serial primary key,
  email text not null unique,
  password    text not null
);

CREATE TABLE IF NOT EXISTS categories (
  id    serial primary key,
  name  text not null unique,
  url   text not null unique
);

CREATE TABLE IF NOT EXISTS products (
  id          serial primary key,
  title       text not null unique,
  description text not null,
  image_url   text not null,
  category_id integer not null,
  featured    boolean
);

CREATE TABLE IF NOT EXISTS product_types (
  id    serial primary key,
  name  text not null unique
);

CREATE TABLE IF NOT EXISTS product_variants (
  id              serial primary key,
  product_id      integer references products(id),
  product_type_id integer references product_types(id),
  price           double precision,
  original_price  double precision
);  

CREATE TABLE IF NOT EXISTS cart (
  id              serial primary key,
  user_id         integer references users(id),
  product_id      integer references products(id),
  produtc_type_id integer references product_types(id),
  quantity        integer
);
