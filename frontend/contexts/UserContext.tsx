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
		hp_url: '',
		location: '',
		github_account_id: '',
		organization_id: '',
		is_deleted: '',
		createdAt: '',
	},
})
const UserContext = () => {
	const [_, setUser] = useRecoilState<User | undefined>(userState)
	setUser(useLogin())
  return <></>
}

export default UserContext
