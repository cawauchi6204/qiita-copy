import axios from 'axios'
import { useEffect, useState } from 'react'
import { User } from '../types/User'

const useLogin = () => {
	const [loginInfo, setLoginInfo] = useState<User>()

	const getUserInfo = async () => {
		const ret = await axios.get<User>('/api/myinfo')
    setLoginInfo(ret.data)
	}
	useEffect(() => {
		getUserInfo()
	}, [])
	return loginInfo
}

export default useLogin
