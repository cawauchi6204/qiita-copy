import { useRef, useState } from "react";
import { useForm } from "react-hook-form"
import Image from "next/image"
import * as S3Client from "../lib/s3Clicent"
import axios from "axios";

type Inputs = {
  file: any
  content: string
};

const ProfilePhotoUpload = () => {
  const [content, setContent] = useState<File>()
  const [fileName, setFileName] = useState("")
  const { register, handleSubmit, setValue, watch } = useForm<Inputs>()
  const inputRef = useRef<HTMLInputElement>(null);

  const onChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    event.preventDefault()
    const { files } = event.target;
    if (!files) return console.error("no file")
    setValue('file', URL.createObjectURL(files[0]));
    setFileName(files[0].name)
    setContent(files[0])
  };
  const { name: previewUrl } = register('file');

  const onSubmit = (data: Inputs) => {
    const uploadFileName = `${data.file.split("blob:http://localhost:3000/")[1]}${fileName}`
    S3Client.upload(uploadFileName, content)
    updateProfileUrl(uploadFileName)
  }

  const updateProfileUrl = async (imgUrl: string) => {
    try {
      return await axios.put("/api/user", {
        imgUrl
      })
    } catch (e) {
      console.error(e)
    }
  }
  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Image src={watch("file")} width={100} height={100} />
      <input {...register('file')} name={previewUrl} ref={inputRef} type="file" onChange={onChange} accept="image/png,image/jpeg" />
      <button className="py-2 px-4 bg-green-500 text-white border-2 border-solid border-green-500">更新する</button>
    </form>
  )
}

export default ProfilePhotoUpload
