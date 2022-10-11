import axios from 'axios'
import { useEffect, useState} from 'react'
import { Userlist } from '../models/user'
//import { Link } from 'react-router-dom'

const UserInfo = (props:any) => {

	const [users, setUsers] = useState(Object)
	// {id:'',username:'',created_at: '',updated_at: ''}

	const [isLoaded, setIsLoaded] = useState(false)
	const [error, setError] = useState("")

	useEffect(() => {
		async function fetchData() {
			const response = await axios.get('users')
			setUsers(response.data)
			setIsLoaded(true)
		}

		fetchData();
		console.log(users)
	}, []);

	if (!isLoaded) {
		return (
			<p>Loading...</p>
		)
	}

	if (error) {	
		return (
			<div>Error: {error}</div>
			)
	} else if (!isLoaded) {
		return (
			<p>Loading...</p>
		)
	}

	console.log(users)

	return (
		<div>
			{JSON.stringify(users)}
      		{users.map((user:any,index:any) => {
			return (
			<div key={index}>
				<h2>name: {user.Name}</h2>
				<h2>作成時刻: {user.Created_at}</h2>
			</div>
        	);
			})};
		</div>
	);	
}

export default UserInfo