import axios from "axios"
import { GetServerSideProps } from "next"
import Link from "next/link"
import { UserPlus } from "react-feather"
import Icon from "../../../../components/Icon"
import Layout from "../../../../components/Layout"
import { Post } from "../../../../types/Post"
import { User } from "../../../../types/User"

type Props = {
  users: User[]
}

const Likes: React.FC<Props> = ({ users }) => {
  return (
    <Layout>
      <div className="flex flex-wrap">
        {
          users.map((user) => {
            return (
              <>
                <div className="flex gap-2">
                  <Icon width={60} height={60} />
                  <div className="flex flex-col">
                    <Link href={`/${user.id}`}>
                      <a>
                        {user.id}
                      </a>
                    </Link>
                    <button className="flex items-center border border-gray-400 border-solid hover:bg-gray-300 p-2">
                      <UserPlus size={16} className="inline" />
                      <span className="text-xs">フォロー</span>
                    </button>
                  </div>
                </div>
              </>
            )
          })
        }
      </div>
    </Layout>
  )
}

export default Likes

export const getServerSideProps: GetServerSideProps = async (context) => {
  if (!context.params) return { props: {} }
  const { userId, postId } = context.params
  const res = await axios.get<Post>(`http://localhost/user/${userId}/items/${postId}/likes`)

  return { props: { users: res.data } }
}
