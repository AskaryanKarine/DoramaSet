interface ErrorMessageProps {
    error: string
}

export function ErrorMessage({ error }: ErrorMessageProps) {
    return (
        <p className="text-center text-xl text-red-600 mt-3">{ error.charAt(0).toUpperCase() + error.slice(1) }</p>
    )
}