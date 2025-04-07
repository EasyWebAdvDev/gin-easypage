# gin-easypage

# Pagination Utils for Gin + GORM

This package includes two utilities for managing **server-side pagination** in web applications built with [Gin](https://github.com/gin-gonic/gin) and [GORM](https://gorm.io/). Both functions allow you to return paginated datasets and include the total number of items in the HTTP response header, which is useful for client-side handling.

---

## Main Features

- Automatic addition of `LIMIT` and `OFFSET` based on the `page` and `page_size` query parameters.
- Counts the total number of rows and sets the `X-Total-Count` header in the HTTP response.
- Supports custom conditions (`WHERE`, `JOIN`) and global search (`globalSearch`).
- Compatible with both GORM queries and raw SQL queries.

---

## Supported URL Parameters

- `page`: the page number to return (0-based, meaning the first page is 0).
- `page_size`: number of items per page.

These parameters must be passed in the HTTP request's query string. For example:

GET /api/users?page=1&page_size=20

---

## `Paginate` Function

This function is designed to be used with standard GORM queries. It allows you to specify:
- The target table.
- Optional conditions (`WHERE`) to apply.
- Custom SQL joins.
- A global search string.

### When to Use

Use this function when performing a classic GORM query on a single table or with declared joins, and you want to paginate the result.

### Generated HTTP Header

In addition to the paginated result, the response will include a header:

X-Total-Count: <total_number_of_results>

This header can be used by the frontend to determine how many pages are available.

---

## `PaginateCustomQuery` Function

This function is intended for use with **raw SQL queries**. It returns a string containing `LIMIT` and `OFFSET`, calculated based on the `page` and `page_size` values in the query string.

### When to Use

Use this function when:
- You have a complex SQL query that cannot easily be expressed with GORM.
- You need to join multiple tables or use subqueries.
- You want to manage the raw SQL directly, while still getting automatic pagination and the `X-Total-Count` header.

### Function Output

Returns a string that can be appended to your SQL query to add pagination, such as:

LIMIT 20 OFFSET 40

In this case as well, the `X-Total-Count` header is automatically added to the HTTP response.

---

## Behavior When `page_size` Is Not Defined

If the `page_size` parameter is missing or set to `0`, the function **does not apply any limit** to the number of results. In this case:
- No `LIMIT/OFFSET` will be applied.
- The full dataset will be returned.

---

## Requirements

- Gin framework
- GORM ORM

---

## Final Notes

The functions are designed to be easily integrated into your project's controllers. For seamless frontend integration, make sure to use the `X-Total-Count`, `page`, and `page_size` values to build dynamic paginated interfaces.

---

## License

This package is released under the MIT license.