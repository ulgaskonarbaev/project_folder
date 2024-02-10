import React, { useState, useEffect } from 'react';
import { List } from 'antd';

function UserList() {
const [users, setUsers] = useState([]);

useEffect(() => {
fetch('/api/users')
.then(response => response.json())
.then(data => setUsers(data));
}, []);

return (
<div>
<h1>User List</h1>
<List
header={<div>List of Users</div>}
bordered
dataSource={users}
renderItem={user => (
<List.Item>
{user.name}
</List.Item>
)}
/>
</div>
);
}

export default UserList;