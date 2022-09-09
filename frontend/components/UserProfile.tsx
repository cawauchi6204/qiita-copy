import Icon from "./Icon"
import Button from "./Button"
import { Hexagon, MapPin } from "react-feather"

type Props = {}

const dummyUserInfo = {
  id: "cawauchi",
  name: "河内",
  description: "エンジニア2年目 2022年の目標 綺麗なコードを書けるようになる ドメイン駆動設計を理解する デザインパターンを習得する UI/UXに関する知識をつける マーケティングの知識を増やし、マーケターとの共通言語を獲得する 説明をうまくする LPICレベル3までとる CCNAとる AWS アソシエイト3冠 体重65kg",
  image: "/vercel",
  hpUrl: "https://note.com/cawauchi",
  location: "TOKYO"
}

const UserProfile: React.FC<Props> = ({ }) => {
  return (
    <div className="w-full bg-white p-6">
      <div className="flex flex-col items-center gap-2 mb-6">
        <Icon width={60} height={60} />
        <p className="font-bold">{dummyUserInfo.name}</p>
        <p className="text-gray-500 text-md">@{dummyUserInfo.id}</p>
        <div className="w-4/5">
          <hr />
          <div className="mt-2 flex text-gray-500 text-xs justify-between text-center">
            <span>31<br />投稿</span>
            <span>8<br />フォロー</span>
            <span>7<br />フォロワー</span>
          </div>
        </div>
      </div>
      <p className="break-all">{dummyUserInfo.description}</p>
      <div className="my-4">
        <Button text="プロフィールを編集する" color="gray" />
      </div>
      <div>
        <Hexagon className="inline mr-2" color="gray" size={15} />
        <span className="text-gray-500 text-xs">{dummyUserInfo.hpUrl}</span>
      </div>
      <div>
        <MapPin className="inline mr-2" color="gray" size={15} />
        <span className="text-gray-500 text-xs">{dummyUserInfo.location}</span>
      </div>
    </div>
  )
}

export default UserProfile
