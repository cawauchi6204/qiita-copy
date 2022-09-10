import { atom, useRecoilState } from 'recoil'
import useLogin from '../hooks/useLogin'
import { User } from '../types/User'

export const userState = atom<User | undefined>({
	key: 'userState',
	default: {
		id: '',
		name: '',
		password: '',
		email: '',
		description: '',
		hpUrl: '',
		location: '',
		githubAccountId: '',
		organizationId: '',
		isDeleted: '',
		createdAt: '',
	},
})
const UserContext = () => {
	const [_, setUser] = useRecoilState<User | undefined>(userState)
	setUser(useLogin())
  return <></>
}

export default UserContext
