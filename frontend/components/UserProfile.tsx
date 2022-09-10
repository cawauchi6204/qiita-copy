import Icon from "./Icon"
import Button from "./Button"
import { Hexagon, MapPin } from "react-feather"
import { User } from "../types/User"

type Props = {
  user: User
}

const UserProfile: React.FC<Props> = ({ user }) => {
  return (
    <div className="w-full bg-white p-6">
      <div className="flex flex-col items-center gap-2 mb-6">
        <Icon width={60} height={60} />
        <p className="font-bold">{user.name}</p>
        <p className="text-gray-500 text-md">@{user.id}</p>
        <div className="w-4/5">
          <hr />
          <div className="mt-2 flex text-gray-500 text-xs justify-between text-center">
            <span>31<br />投稿</span>
            <span>8<br />フォロー</span>
            <span>7<br />フォロワー</span>
          </div>
        </div>
      </div>
      <p className="break-all">{user.description}</p>
      <div className="my-4">
        <Button text="プロフィールを編集する" color="gray" />
      </div>
      <div>
        <Hexagon className="inline mr-2" color="gray" size={15} />
        <span className="text-gray-500 text-xs">{user.hpUrl}</span>
      </div>
      <div>
        <MapPin className="inline mr-2" color="gray" size={15} />
        <span className="text-gray-500 text-xs">{user.location}</span>
      </div>
    </div>
  )
}

export default UserProfile
