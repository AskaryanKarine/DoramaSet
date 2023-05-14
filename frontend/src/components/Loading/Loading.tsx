export function Loading() {
    return (
        <>
            <div className="absolute bg-black/50 top-0 right-0 left-0 bottom-0 h-screen flex justify-center items-center">
                <div className='p-5 rounded-2xl bg-white'>
                    <i className="fa-solid fa-spinner fa-spin-pulse"></i>
                </div>
            </div>
        </>
    )
}