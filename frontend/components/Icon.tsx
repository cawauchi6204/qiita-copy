import Image from "next/image"

type Props = {
  src?: string
  width?: number
  height?: number
}

const Icon: React.FC<Props> = ({ src = "/tori.png", width = 40, height = 40 }) => {
  return (
    <div style={{ "width": width, "height": height }}>
      <Image className="rounded-full" src={src} width={width} height={height} />
    </div>
  )
}

export default Icon
