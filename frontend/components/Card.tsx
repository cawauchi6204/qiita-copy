import Link from "next/link"
import { User, Pocket, Heart, Tag } from "react-feather"
import { Post } from "../types/Post"

const Card: React.FC<Pick<Post, "id" | "title" | "postedBy" | "createdAt">> = ({ id, title, postedBy, createdAt }) => {
  return (
    <div className="w-full bg-white flex flex-col p-4">
      <div className="flex">
        <User size={40} />
        <div className="flex flex-col">
          <Link href={`/${postedBy}`}>
            <a className="hover:underline inline-block">
              {postedBy}
            </a>
          </Link>
          <time className="text-xs text-gray">
            {createdAt}
          </time>
        </div>
      </div>
      <div className="flex flex-col ml-10">
        <Link href={`/${postedBy}/items/${id}`}>
          <a className="text-2xl mt-2 font-bold hover:underline">
            {title}
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
