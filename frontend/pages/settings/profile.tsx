import axios from "axios";
import { useForm, SubmitHandler } from "react-hook-form";
import { useRecoilValue } from "recoil";

import Icon from "../../components/Icon"
import Layout from "../../components/Layout"
import SettingBar from "../../components/SettingBar"
import { userState } from "../../contexts/UserContext";

type Inputs = {
  name: string
  email: string
  hpUrl: string
  organization: string
  location: string
  description: string
};

type Props = {}

const Profile: React.FC<Props> = ({ }) => {
  const {
    register,
    handleSubmit,
  } = useForm<Inputs>();
  const user = useRecoilValue(userState);
  if (!user) return <></>

  const onSubmit: SubmitHandler<Inputs> = async (i) => {
    const data = { ...i, id: user.id }
    await axios.put("/api/user", data)
    console.log("onSubmit:", data)
  };

  return (
    <Layout className="flex gap-4 max-w-6xl m-auto">
      <div className="w-1/4">
        <SettingBar />
      </div>
      <div className="w-3/4 bg-white p-8">
        <form onSubmit={handleSubmit(onSubmit)}>
          <div className="flex justify-center items-center">
            <Icon width={30} height={30} /><span className="text-2xl">　cawauchi　アカウント　/　</span><span className="text-2xl font-bold">公開用プロフィール</span>
          </div>
          <div className="mt-8">
            <p>名前</p>
            <input defaultValue={user.name} type="text" {...register("name")} className="w-full border border-solid border-gray p-2" />
          </div>
          <div className="mt-8">
            <p>公開メールアドレス</p>
            <input type="checkbox" name="" id="" />メールアドレスを全体に公開する
            <input defaultValue={user.email} type="email" {...register("email")} className="w-full border border-solid border-gray p-2" />
          </div>
          <div className="mt-8">
            <p>サイト/ブログ</p>
            <input defaultValue={user.hpUrl} type="text" {...register("hpUrl")} className="w-full border border-solid border-gray p-2" />
          </div>
          <div className="mt-8">
            <p>所属している組織・企業</p>
            <input type="text" {...register("organization")} className="w-full border border-solid border-gray p-2" />
          </div>
          <div className="mt-8">
            <p>居住地</p>
            <input defaultValue={user.location} type="text" {...register("location")} className="w-full border border-solid border-gray p-2" />
          </div>
          <div className="mt-8">
            <p>自己紹介</p>
            <textarea defaultValue={user.description} {...register("description")} className="w-full border border-solid border-gray p-2" />
          </div>
          <button type="submit" className="mt-6 text-white px-6 py-2 text-xs" style={{ "background": "#55C500" }}>更新する</button>
        </form>
      </div>
    </Layout>
  )
}

export default Profile
