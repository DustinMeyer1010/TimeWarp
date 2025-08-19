📘 TimeWarp API Routes

This document describes all available API routes in TimeWarp, including their purpose, required parameters, and return values.

🔑 Accounts

Routes related to account creation, authentication, and management.

|   Method   | Path             | Description             | Requirements                    | Returns               |
| :--------: | ---------------- | ----------------------- | ------------------------------- | --------------------- |
|  **POST**  | `/account`       | Create a new account    | `email`, `username`, `password` | —                     |
|  **POST**  | `/account/login` | Log in to an account    | `username`, `password`          | Refresh & Auth tokens |
|   **GET**  | `/refresh`       | Refresh an auth token   | —                               | New auth token        |
| **DELETE** | `/account/{id}`  | Delete an account by ID | `account_id`                    | —                     |



📅 Habits

Habit routes are secured with an auth token. You cannot access or modify another user’s habits.
If the provided account_id or habit_id does not match the authenticated user, no data will be returned and no changes will be made.

|   Method   | Path                   | Description                   | Requirements | Returns                         |
| :--------: | ---------------------- | ----------------------------- | ------------ | ------------------------------- |
|   **GET**  | `/account/habits/{id}` | Get all habits for an account | `account_id` | List of habits for that account |
|   **GET**  | `/habit/{id}`          | Get a single habit            | `habit_id`   | Habit object                    |
| **DELETE** | `/habit/{id}`          | Delete a habit                | `habit_id`   | —                               |


