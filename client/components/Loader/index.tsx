import React from "react";

type PropTypes = {
    loading: boolean;
    message?: string;
    children?: React.ReactNode;
};

const Loader = ({ loading, message, children }: PropTypes) => {
    return <div>{loading ? <div>{message}</div> : children}</div>;
};

export default Loader;
