import Link from "next/link"
import { User, Pocket, Heart, Tag } from "react-feather"

type Props = {}

const Card: React.FC<Props> = ({ }) => {
  return (
    <div className="w-full bg-white flex flex-col p-4">
      <div className="flex">
        <User size={40} />
        <div className="flex flex-col">
          <Link href="">
            <a>
              @hogehoge(Hoge Taro)
            </a>
          </Link>
          <time className="text-xs text-gray">
            2022年09月06日
          </time>
        </div>
      </div>
      <div className="flex flex-col ml-10">
        <Link href="">
          <a className="text-2xl mt-2 font-bold hover:underline">
            初学者が覚えたいチーム開発でのGit操作
          </a>
        </Link>
      </div>
      <div className="ml-10">
        <Tag className="inline" size={16} />
        <Link href="">
          <a>
            Git
          </a>
        </Link>
        <div className="flex justify-between mt-2">
          <Heart size={16} />
          <Pocket size={40} />
        </div>
      </div>
    </div >
  )
}

export default Card
