import axios from "axios"
import { GetServerSideProps } from "next"
import { useRecoilValue } from "recoil"
import CardList from "../../components/CardList"
import Layout from "../../components/Layout"
import UserProfile from "../../components/UserProfile"
import { userState } from "../../contexts/UserContext"
import { Post } from "../../types/Post"

type Props = {
  posts: Post[]
}

const UserId: React.FC<Props> = ({ posts }) => {
  const user = useRecoilValue(userState);
  if (!user) return <></>
  return (
    <Layout>
      <div className="max-w-5xl m-auto">
        <div className="flex gap-4">
          <div className="w-1/3">
            <UserProfile user={user} />
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
  if (!context.params) return { props: {} }
  const { userId } = context.params;
  const posts = await axios.get(`http://localhost/user/${userId}/posts`)
  return { props: { posts: posts.data } }
}

export default UserId
