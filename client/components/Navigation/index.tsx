import Link from "next/link";

type PropTypes = {
    brandName: string;
};

const Navigation = ({ brandName }: PropTypes) => {
    return (
        <nav>
            <div>
                <Link href="/">{brandName}</Link>
            </div>
        </nav>
    );
};

export default Navigation;
