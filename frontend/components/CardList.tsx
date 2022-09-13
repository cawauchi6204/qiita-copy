import Card from "./Card"
import { Post } from "../types/Post"

type Props = {
  posts: Post[]
}

const CardList: React.FC<Props> = ({ posts }) => {
  return (
    < div className="flex flex-col gap-4" >
      {
        posts.map((post) => {
          return (
            <Card key={post.id} post={post} />
          )
        })
      }
    </div >
  )
}

export default CardList
