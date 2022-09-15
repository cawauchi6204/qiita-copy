import { useRecoilValue } from "recoil";

import Icon from "../../components/Icon";
import Layout from "../../components/Layout"
import ProfilePhotoUpload from "../../components/ProfilePhotoUpload"
import SettingBar from "../../components/SettingBar"
import { userState } from "../../contexts/UserContext";

type Props = {}

const customImage: React.FC<Props> = ({ }) => {
  const user = useRecoilValue(userState);
  if (!user) return <></>
  return (
    <Layout className="flex gap-4">
      <div className="w-1/4">
        <SettingBar />
      </div>
      <div className="w-3/4 bg-white p-8">
        <div className="flex justify-center items-center">
          <Icon width={30} height={30} /><span className="text-2xl">　cawauchi　アカウント　/　</span><span className="text-2xl font-bold">プロフィール画像アップロード</span>
        </div>
        <div className="mt-20">
          <ProfilePhotoUpload />
        </div>
      </div>
    </Layout>
  )
}

export default customImage
