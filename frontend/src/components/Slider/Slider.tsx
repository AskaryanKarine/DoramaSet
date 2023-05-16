import {IPhoto} from "../../models/IPhoto";
import SimpleImageSlider from "react-simple-image-slider";

interface SliderProps {
    photo?: IPhoto[]
}

export function Slider({photo}:SliderProps) {
    return (<>
        <div className="w-auto mr-5 relative">
            {photo ? <SimpleImageSlider
                width={300}
                height={400}
                autoPlay={true}
                images={photo}
                showBullets={false}
                showNavs={true}
                navSize={20}
            /> : <i className="fa-regular fa-image fa-5x" style={{color: "#787d87"}}></i>}
        </div>
    </>)
}