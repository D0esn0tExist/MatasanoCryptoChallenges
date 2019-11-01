def conv_hex_base64():
    # converts hex string to base64
    
    base64 = string.decode("hex").encode("base64")
    print(base64)
    return base64

if __name__ == "__main__":
    string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
    conv_hex_base64()