import {toUpperFirst} from "../../hooks/toUpperFirst";

interface ErrorMessageProps {
    error: string
    info?: boolean
}

export function ErrorMessage({ error }: ErrorMessageProps) {
    return (
        <p className="text-center text-xl text-red-600 mt-3">{ toUpperFirst(error) }</p>
    )
}

export function ErrorMessage1({ error, info }: ErrorMessageProps) {
    const icon = info ? "fa-info" : "fa-exclamation"
    // <i className="fa-solid fa-info"></i> - info
    // <i className="fa-solid "></i> - warning
    return (
        <p className="text-center text-xl text-red-600 mt-3">{ toUpperFirst(error) }</p>
    )
}