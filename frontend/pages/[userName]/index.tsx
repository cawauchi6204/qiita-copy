import axios from "axios"
import { GetServerSideProps } from "next"
import CardList from "../../components/CardList"
import Layout from "../../components/Layout"
import UserProfile from "../../components/UserProfile"
import { Post } from "../../types/Post"

type Props = {
  posts: Post[]
}

const PostId: React.FC<Props> = ({ posts }) => {
  return (
    <Layout>
      <div className="max-w-5xl m-auto">
        <div className="flex gap-4">
          <div className="w-1/3">
            <UserProfile />
          </div>
          <div className="w-2/3">
            <CardList posts={posts} />
          </div>
        </div>
      </div>
    </Layout>
  )
}

export const getServerSideProps: GetServerSideProps = async (context) => {
  const { userName } = context.query;
  const res = await axios.get(`http://localhost/users/${userName}/posts`)

  return { props: { posts: res.data } }
}

export default PostId
