import Card from "./Card"
import { Post } from "../types/Post"

type Props = {
  posts: Pick<Post, "id" | "title" | "createdAt">[]
}

const CardList: React.FC<Props> = ({ posts }) => {
  return (
    < div className="flex flex-col gap-4" >
      {
        posts.map((post) => {
          return (
            <Card id={post.id} title={post.title} createdAt={post.createdAt} />
          )
        })
      }
    </div >
  )
}

export default CardList
